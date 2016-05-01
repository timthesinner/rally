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

var rally, ERROR = NewRally("", "")

func TestRallyLogin(t *testing.T) {
	fmt.Println("Starting test")
	if ERROR != nil {
		t.Fatal(ERROR)
	} else {
		rally.LogRequests = true
		if features, err := rally.RawGet("portfolioitem/feature?fetch=true&query=(Name%20%3D%20test)"); err != nil {
			t.Fatal(err)
		} else {
			fmt.Println(len(features.Results))

			if f, _ := os.Create(path.Join("./", "Features.json")); err == nil {
				defer f.Close()
				if b, err := json.MarshalIndent(features, "", "  "); err == nil {
					f.Write(b)
				}
			}
		}
	}
}

func TestFeatures(t *testing.T) {
	if ERROR != nil {
		t.Fatal(ERROR)
	} else {
		features, _ := rally.Features("(FormattedID%20%3D%20Test)")
		fmt.Println(PrettyJSON(features))
	}
}
