// Copyright 2022
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package conf

import (
	"github.com/taouniverse/tao"
)

func init() {
	data := []byte(`
{
    "tao": {
		"log": {
		  "level": "debug"
		},
    	"hide_banner": true
  	},
    "hello": {
        "print": "this is code config!",
        "times": 2
    }
}
`)
	err := tao.SetConfigBytesAll(data, tao.Json)
	if err != nil {
		tao.Panic(err)
	}
}
