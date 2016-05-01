//Package rally by TimTheSinner
package rally

/**
 * Copyright (c) 2016 TimTheSinner All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"
)

//GetMap Return a map for a key
func GetMap(key string, m map[string]interface{}) (map[string]interface{}, bool) {
	if m != nil {
		if value, ok := m[key]; ok {
			v, ok := value.(map[string]interface{})
			return v, ok
		}
	}
	return nil, false
}

//GetArray Return an array for a key
func GetArray(key string, m map[string]interface{}) ([]interface{}, bool) {
	if m != nil {
		if value, ok := m[key]; ok {
			v, ok := value.([]interface{})
			return v, ok
		}
	}
	return nil, false
}

//PrettyJSON pretty print JSON
func PrettyJSON(obj interface{}) (string, error) {
	pretty, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}

type mapping struct {
	src  string
	targ string
}

var mappings = []mapping{
	mapping{src: "_ref", targ: "URI"},
	mapping{src: "_refObjectName", targ: "RefObjectName"},
	mapping{src: "_rallyAPIMinor", targ: "APIMinor"},
	mapping{src: "_rallyAPIMajor", targ: "APIMajor"},
	mapping{src: "_refObjectUUID", targ: "RefObjectUUID"},
	mapping{src: "_tagsNameArray", targ: "Names"},
	mapping{src: "_CreatedAt", targ: "CreatedAt"},
	mapping{src: "_objectVersion", targ: "ObjectVersion"},
	mapping{src: "_type", targ: "RefType"},
	mapping{src: "VersionId", targ: "VersionID"},
}

//Clean strip out annoying tags
func Clean(o interface{}) {
	if m, ok := o.(map[string]interface{}); ok {
		for _, mapper := range mappings {
			if src, ok := m[mapper.src]; ok {
				m[mapper.targ] = src
			}
		}

		for k, v := range m {
			if strings.HasPrefix(k, "_") {
				delete(m, k)
			} else if m2, ok := v.(map[string]interface{}); ok {
				Clean(m2)
			} else if a, ok := v.([]interface{}); ok {
				Clean(a)
			}
		}
	} else if a, ok := o.([]interface{}); ok {
		for _, v := range a {
			Clean(v)
		}
	}
}

//Client struct enabling authenticated access to Rally
type Client struct {
	jar           *cookiejar.Jar
	client        *http.Client
	Server        string
	APIVersion    string
	SecurityToken string
	Workspace     string
	Workspaces    map[string]string
	AlwaysFetch   bool
	LogRequests   bool
}

//PaginatedGet get all results
func (r *Client) PaginatedGet(start float64, path string) (*QueryResult, error) {
	return r.paginatedGet(start, -1, path, nil)
}

//RawGet run a custom query (any string)
func (r *Client) RawGet(raw string) (*QueryResult, error) {
	return r.PaginatedGet(1, raw)
}

//Path normalize a URI request, if the URI is missing https the client server and version are inserted
func (r *Client) Path(path string) string {
	if strings.HasPrefix(path, "https:") {
		return path
	}
	return fmt.Sprintf("%s/slm/webservice/%s/%s", r.Server, r.APIVersion, path)
}

//Features list a feature for query
func (r *Client) Features(query string) (features []PortfolioItemFeature, err error) {
	result, err := r.PaginatedGet(1, "portfolioitem/feature?fetch=true&query="+query)
	if err != nil {
		return features, err
	}

	for _, res := range result.Results {
		var buff bytes.Buffer
		json.NewEncoder(&buff).Encode(res)

		var feature PortfolioItemFeature
		if err := json.NewDecoder(&buff).Decode(&feature); err != nil {
			fmt.Println(err)
			fmt.Println(PrettyJSON(res))
			return features, err
		}

		features = append(features, feature)
	}

	return features, nil
}

//Login user credentials are used to authenitcate with Rally, this is necessary only once.
func (r *Client) Login(username, password string) error {
	if username == "" {
		username = os.Getenv("RALLY_USERNAME")
	}

	if password == "" {
		password = os.Getenv("RALLY_PASSWORD")
		if decoded, err := base64.StdEncoding.DecodeString(password); err == nil {
			password = string(decoded)
		}
	}

	req, err := http.NewRequest("GET", r.Path("security/authorize"), nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(username, password)
	res, err := r.client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return errors.New(res.Status)
	}

	var body map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return err
	}

	if result, ok := GetMap("OperationResult", body); ok {
		if r.SecurityToken, ok = result["SecurityToken"].(string); !ok {
			return errors.New("Could not find SecurityToken in result")
		}
		return nil
	}

	rJSON, err := PrettyJSON(body)
	if err != nil {
		return err
	}
	return errors.New("Could not find security token in" + rJSON)
}

//InitWorkspaces initialize all workspaces available to the user
func (r *Client) InitWorkspaces() {
	if result, err := r.PaginatedGet(1, "subscription?fetch=true&query="); err == nil && len(result.Results) > 0 {
		if rWorkspaces, ok := GetMap("Workspaces", result.Results[0]); ok {
			if workspaces, err := r.PaginatedGet(1, rWorkspaces["_ref"].(string)+"?fetch=true"); err == nil && len(workspaces.Results) > 0 {
				r.Workspaces = make(map[string]string)
				for _, result := range workspaces.Results {
					if r.Workspace == "" {
						r.Workspace = result["_ref"].(string)
					}
					r.Workspaces[result["Name"].(string)] = result["_ref"].(string)
				}
			}
		}
	}
}

//NewRally initialize a rally client
func NewRally(username, password string) (*Client, error) {
	var rallyJar, _ = cookiejar.New(nil)

	client := &Client{
		jar:        rallyJar,
		client:     &http.Client{Jar: rallyJar, Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: false}}},
		Server:     "https://rally1.rallydev.com",
		APIVersion: "v2.0",
		Workspace:  os.Getenv("RALLY_WORKSPACE"),
	}

	err := client.Login(username, password)
	if err == nil {
		client.InitWorkspaces()
	}

	return client, err

}
