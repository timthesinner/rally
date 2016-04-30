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
package rally

import (
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

func GetMap(key string, m map[string]interface{}) (map[string]interface{}, bool) {
	if m != nil {
		if value, ok := m[key]; ok {
			v, ok := value.(map[string]interface{})
			return v, ok
		}
	}
	return nil, false
}

func GetArray(key string, m map[string]interface{}) ([]interface{}, bool) {
	if m != nil {
		if value, ok := m[key]; ok {
			v, ok := value.([]interface{})
			return v, ok
		}
	}
	return nil, false
}

func PrettyJson(obj interface{}) (string, error) {
	if pretty, err := json.MarshalIndent(obj, "", "  "); err != nil {
		return "", err
	} else {
		return string(pretty), nil
	}
}

type mapping struct {
	src  string
	targ string
}

var mappings = []mapping{
	mapping{src: "_ref", targ: "URI"},
	mapping{src: "_refObjectName", targ: "ObjectName"},
	mapping{src: "_refObjectUUID", targ: "UUID"},
	mapping{src: "_tagsNameArray", targ: "Names"},
	mapping{src: "_type", targ: "Type"},
}

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

type RallyClient struct {
	jar           *cookiejar.Jar
	client        *http.Client
	Server        string
	ApiVersion    string
	SecurityToken string
	Workspace     string
	Workspaces    map[string]string
	AlwaysFetch   bool
}

func (r *RallyClient) PaginatedGet(start float64, path string) (map[string]interface{}, error) {
	return r.paginatedGet(start, -1, path, nil)
}

func (r *RallyClient) RawGet(raw string) (body map[string]interface{}, err error) {
	return r.PaginatedGet(1, raw)
}

func (r *RallyClient) Path(path string) string {
	if strings.HasPrefix(path, "https:") {
		return path
	}
	return fmt.Sprintf("%s/slm/webservice/%s/%s", r.Server, r.ApiVersion, path)
}

func (r *RallyClient) Login(username, password string) error {
	if username == "" {
		username = os.Getenv("RALLY_USERNAME")
	}

	if password == "" {
		password = os.Getenv("RALLY_PASSWORD")
		if decoded, err := base64.StdEncoding.DecodeString(password); err == nil {
			password = string(decoded)
		}
	}

	if req, err := http.NewRequest("GET", r.Path("security/authorize"), nil); err != nil {
		return err
	} else {
		req.SetBasicAuth(username, password)
		if res, err := r.client.Do(req); err != nil {
			return err
		} else {
			defer res.Body.Close()
			if res.StatusCode != 200 {
				return errors.New(res.Status)
			}

			var body map[string]interface{}
			if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
				return err
			}

			if result, ok := GetMap("OperationResult", body); ok {
				if r.SecurityToken, ok = result["SecurityToken"].(string); !ok {
					return errors.New("Could not find SecurityToken in result")
				}
			} else if rJson, err := PrettyJson(body); err != nil {
				return err
			} else {
				return errors.New("Could not find security token in" + rJson)
			}
		}
	}

	return nil
}

func (r *RallyClient) InitWorkspaces() {
	if raw, err := r.PaginatedGet(1, "subscription?fetch=true&query="); err == nil {
		if QueryResult, ok := GetMap("QueryResult", raw); ok {
			if Results, ok := GetArray("Results", QueryResult); ok && len(Results) == 1 {
				if Workspaces, ok := GetMap("Workspaces", Results[0].(map[string]interface{})); ok {
					if Workspaces, err := r.PaginatedGet(1, Workspaces["_ref"].(string)+"?fetch=true"); err == nil {
						if QueryResult, ok := GetMap("QueryResult", Workspaces); ok {
							if Results, ok := GetArray("Results", QueryResult); ok && len(Results) != 0 {
								r.Workspaces = make(map[string]string)
								for _, result := range Results {
									if result, ok := result.(map[string]interface{}); ok {
										if r.Workspace == "" {
											r.Workspace = result["_ref"].(string)
										}
										r.Workspaces[result["Name"].(string)] = result["_ref"].(string)
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func NewRally(username, password string) (*RallyClient, error) {
	var rallyJar, _ = cookiejar.New(nil)

	client := &RallyClient{
		jar:        rallyJar,
		client:     &http.Client{Jar: rallyJar, Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: false}}},
		Server:     "https://rally1.rallydev.com",
		ApiVersion: "v2.0",
		Workspace:  os.Getenv("RALLY_WORKSPACE"),
	}

	err := client.Login(username, password)
	if err == nil {
		client.InitWorkspaces()
	}

	return client, err
}
