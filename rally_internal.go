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
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

func (r *Client) get(start float64, path string) (map[string]interface{}, error) {
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
	//fmt.Println("GET", URI)

	res, err := r.client.Get(URI)
	if err != nil {
		return make(map[string]interface{}), err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return make(map[string]interface{}), errors.New(res.Status)
	}

	var parse map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&parse); err != nil {
		return make(map[string]interface{}), err
	}
	return parse, nil
}

func (r *Client) paginatedGet(start, total float64, path string, body map[string]interface{}) (map[string]interface{}, error) {
	if total != -1 && start > total {
		return body, nil
	}

	parse, err := r.get(start, path)
	if err != nil {
		return body, err
	}
	if result, ok := GetMap("QueryResult", parse); ok {
		if body == nil {
			return r.paginatedGet(start+20, result["TotalResultCount"].(float64), path, parse)
		} else if bResult, ok := GetMap("QueryResult", body); ok {
			if results, ok := GetArray("Results", result); ok {
				if bResults, ok := GetArray("Results", bResult); ok {
					bResult["Results"] = append(bResults, results...)
				}
			}
			return r.paginatedGet(start+20, bResult["TotalResultCount"].(float64), path, body)
		} else {
			return body, nil
		}
	} else if body == nil {
		return parse, nil
	} else {
		return body, nil
	}
}
