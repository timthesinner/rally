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
	"fmt"
	"os"
	"path"
	"sort"
	"strings"
)

func (r *Client) processModel() {
	r.AlwaysFetch = true
	types := make(map[string]*map[string]interface{})

	if rSubscription, err := r.getRaw(1, "subscription"); err == nil {
		fmt.Println(PrettyJSON(rSubscription))
		if Subscription, ok := GetMap("Subscription", rSubscription); ok {
			types["Subscription"] = &Subscription
			if RevisionHistory, ok := GetMap("RevisionHistory", Subscription); ok {
				r.processModelQuery(types, RevisionHistory["_ref"].(string))
			}
			if Workspaces, ok := GetMap("Workspaces", Subscription); ok {
				r.processModelQuery(types, Workspaces["_ref"].(string))
			}
			r.processModelQuery(types, "portfolioitem/feature")
			r.processModelQuery(types, "portfolioitem/theme")
		}
	}

	var orderedType []string
	for k := range types {
		orderedType = append(orderedType, k)
	}
	sort.Strings(orderedType)

	golangTypes := []string{`//Package rally by TimTheSinner (this file was auto generated)
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

//Base minimal Rally object type
type Base struct {
 	APIMajor         string                 ` + "`json:\"_rallyAPIMajor\"`" + `
 	APIMinor         string                 ` + "`json:\"_rallyAPIMinor\"`" + `
 	URI              string                 ` + "`json:\"_ref\"`" + `
 	RefObjectName    string                 ` + "`json:\"_refObjectName\"`" + `
 	RefObjectUUID    string                 ` + "`json:\"_refObjectUUID\"`" + `
 	RefType          string                 ` + "`json:\"_type\"`" + `
}

//QueryResult standard query result
type QueryResult struct {
	Errors           []interface{}
	PageSize         float64
	Results          []interface{}
  StartIndex       float64
  TotalResultCount float64
  Warnings         []interface{}
 	APIMajor         string                 ` + "`json:\"_rallyAPIMajor\"`" + `
 	APIMinor         string                 ` + "`json:\"_rallyAPIMinor\"`" + `
}`}

	for _, k := range orderedType {
		fmt.Println(k)
		vp := types[k]
		v := *vp
		k = strings.Replace(k, "/", "", 1)
		gType := fmt.Sprintf("//%v auto-generated\ntype %v struct {\n", k, k)
		attrs := make([]string, len(v))
		for attr := range v {
			attrs = append(attrs, attr)
		}
		sort.Strings(attrs)

		for _, attr := range attrs {
			if strings.HasPrefix(attr, "c_") {
				continue
			}

			val := v[attr]
			gAttr := attr
			for _, mapper := range mappings {
				if attr == mapper.src {
					gAttr = mapper.targ
					break
				}
			}

			switch v := val.(type) {
			case int, float64:
				gType += fmt.Sprintf("  %v  float64  `json:\"%v\"`\n", gAttr, attr)
			case string:
				gType += fmt.Sprintf("  %v  string  `json:\"%v\"`\n", gAttr, attr)
			case []interface{}:
				gType += fmt.Sprintf("  %v  []interface{}  `json:\"%v\"`\n", gAttr, attr)
			case map[string]interface{}:
				if Type, ok := v["_type"]; ok {
					Type := strings.Replace(Type.(string), "/", "", 1)
					isArray := false
					if _, ok := v["Count"]; ok {
						isArray = true
					}

					if (!isArray && k == Type) || (k == "Workspace" && Type == "WorkspaceConfiguration") || (k == "RevisionHistory" && Type == "Subscription") {
						gType += fmt.Sprintf("  %v  Base  `json:\"%v\"`\n", gAttr, attr)
						break
					} else {
						if _, ok := types[Type]; ok {
							if isArray {
								Type = "[]" + Type
							}
							gType += fmt.Sprintf("  %v  %v  `json:\"%v\"`\n", gAttr, Type, attr)
							break
						}
					}
				}
				gType += fmt.Sprintf("  %v  map[string]interface{}  `json:\"%v\"`\n", gAttr, attr)
			}
		}
		gType += "  CustomAttributes  map[string]interface{}  `json:\"CustomAttributes\"`\n"
		gType += "}"
		golangTypes = append(golangTypes, gType)
	}

	if f, err := os.Create(path.Join("./", "Model.json")); err == nil {
		defer f.Close()
		if b, err := json.MarshalIndent(types, "", "  "); err == nil {
			f.Write(b)
		}
	}

	if f, err := os.Create(path.Join("./", "rally_objects.go")); err == nil {
		defer f.Close()
		f.WriteString(strings.Join(golangTypes, "\n\n"))
	}
}

func firstKey(m map[string]interface{}) string {
	if len(m) == 1 {
		for k := range m {
			return k
		}
	}
	return ""
}

func (r *Client) processModelQuery(types map[string]*map[string]interface{}, url string) {
	recurse := func(obj interface{}) bool {
		if obj, ok := obj.(map[string]interface{}); ok {
			if Type, ok := obj["_type"]; ok {
				if _, ok := types[Type.(string)]; !ok {
					r.processModelQuery(types, obj["_ref"].(string))
				}
			}
			return true
		}
		return false
	}

	process := func(Result map[string]interface{}) {
		if Type, ok := Result["_type"]; ok {
			if _, ok := types[Type.(string)]; !ok {
				types[Type.(string)] = &Result
			}
		} else {
			fmt.Println(PrettyJSON(Result))
		}

		for _, v := range Result {
			if ok := recurse(v); ok {
				continue
			} else if a, ok := v.([]interface{}); ok && len(a) > 0 {
				recurse(a[0])
			}
		}
	}

	if raw, err := r.getRaw(1, url); err == nil {
		if QueryResult, ok := GetMap("QueryResult", raw); !ok {
			Type := firstKey(raw)
			if Result, ok := GetMap(Type, raw); ok {
				Result["_type"] = Type
				process(Result)
			}
		} else {
			if Results, ok := GetArray("Results", QueryResult); ok && len(Results) > 0 {
				for _, Result := range Results {
					process(Result.(map[string]interface{}))
				}
			}
		}
	}
}
