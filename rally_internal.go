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
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

func (r *RallyClient) paginatedGet(start, total float64, path string, body map[string]interface{}) (map[string]interface{}, error) {
	if total != -1 && start > total {
		return body, nil
	}

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

	URI := r.Path(path + sep + fmt.Sprintf("start=%v&count=%v%v%v", start, 20, fetch, workspace))
	fmt.Println("GET", URI)
	if res, err := r.client.Get(URI); err != nil {
		return body, err
	} else {
		defer res.Body.Close()
		if res.StatusCode != 200 {
			return body, errors.New(res.Status)
		}

		return func(body map[string]interface{}) (map[string]interface{}, error) {
			var parse map[string]interface{}
			if err := json.NewDecoder(res.Body).Decode(&parse); err != nil {
				return body, err
			}

			//Clean(parse)
			//fmt.Println(PrettyJson(parse))
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
		}(body)
	}
}

func (r *RallyClient) processModel() {
	r.AlwaysFetch = true
	types := make(map[string]map[string]interface{})

	if Subscription, err := r.paginatedGet(1, -1, "subscription", nil); err == nil {
		if Subscription, ok := GetMap("Subscription", Subscription); ok {
			types["Subscription"] = Subscription
			if RevisionHistory, ok := GetMap("RevisionHistory", Subscription); ok {
				r.processModelQuery(types, RevisionHistory["_ref"].(string))
			}
			if Workspaces, ok := GetMap("Workspaces", Subscription); ok {
				r.processModelQuery(types, Workspaces["_ref"].(string))
			}
			r.processModelQuery(types, "portfolioitem/feature")
		}
	}

	for k, _ := range types {
		fmt.Println(k)
	}

	fmt.Println(types)
}

func (r *RallyClient) processModelQuery(types map[string]map[string]interface{}, url string) {
	if raw, err := r.paginatedGet(1, -1, url, nil); err == nil {
		if QueryResult, ok := GetMap("QueryResult", raw); ok {
			if Results, ok := GetArray("Results", QueryResult); ok && len(Results) != 0 {
				Result := Results[0].(map[string]interface{})
				if Type, ok := Result["_type"]; ok {
					types[Type.(string)] = Result
				} else {
					fmt.Println(PrettyJson(Result))
				}

				for k, v := range Result {
					if obj, ok := v.(map[string]interface{}); ok {
						if Type, ok := obj["_type"]; ok {
							if _, ok := types[Type.(string)]; !ok {
								r.processModelQuery(types, obj["_ref"].(string))
							}
						}
					} else if _, ok := v.([]interface{}); ok {
						fmt.Println(k)
					}
				}
			}
		}
	}
}
