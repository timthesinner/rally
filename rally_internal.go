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
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

func (r *Client) getRaw(start float64, path string) (raw map[string]interface{}, err error) {
	sep := "?"
	if strings.Contains(path, "?") {
		sep = "&"
	}

	fetch := ""
	if (r.AlwaysFetch) && !strings.Contains(path, "fetch") {
		fetch = "&fetch=true"
	}

	workspace := ""
	if r.Workspace != "" && !strings.Contains(path, "workspace") {
		workspace = "&workspace=" + r.Workspace
	}

	if start < 1 {
		start = 1
	}

	URI := r.Path(path + sep + fmt.Sprintf("start=%v&count=%v%v%v", start, 20, fetch, workspace))
	if r.LogRequests {
		fmt.Println("GET", URI)
	}

	res, err := r.client.Get(URI)
	if err != nil {
		return raw, err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return raw, errors.New(res.Status)
	}

	err = json.NewDecoder(res.Body).Decode(&raw)
	return raw, err
}

func (r *Client) get(start float64, path string) (results *QueryResult, err error) {
	raw, err := r.getRaw(start, path)
	if err != nil {
		return results, err
	}

	if res, ok := raw["QueryResult"]; ok {
		var buff bytes.Buffer
		json.NewEncoder(&buff).Encode(res)
		err = json.NewDecoder(&buff).Decode(&results)
		return results, err
	}
	return results, errors.New("Result did not have a QueryResult field")
}

func (r *Client) paginatedGet(start, total float64, path string, body *QueryResult) (*QueryResult, error) {
	if total != -1 && start > total {
		return body, nil
	}

	parse, err := r.get(start, path)
	if err != nil {
		return body, err
	}

	if body == nil {
		return r.paginatedGet(start+20, parse.TotalResultCount, path, parse)
	} else if len(body.Results) > 0 {
		body.Results = append(body.Results, parse.Results...)
	}
	return r.paginatedGet(start+20, parse.TotalResultCount, path, body)
}
