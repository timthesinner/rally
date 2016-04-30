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
	"fmt"
	"os"
	"path"
	"testing"
)

func TestRallyLogin(t *testing.T) {
	if rally, err := NewRally("", ""); err != nil {
		t.Fatal(err)
	} else {
		body, err := rally.RawGet("portfolioitem/feature?fetch=true")
		q, _ := GetMap("QueryResult", body)
		if results, ok := GetArray("Results", q); ok {
			fmt.Println(len(results))
		}

		rally.processModel()

		if f, _ := os.Create(path.Join("./", "Features.json")); err == nil {
			defer f.Close()
			if b, err := json.MarshalIndent(body, "", "  "); err == nil {
				f.Write(b)
			}
		}
	}
}
