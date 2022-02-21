// Copyright 2021
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
package main

import (
	//_ "github.com/taouniverse/hello/conf" // 1. code config
	"github.com/taouniverse/tao" // 0. defaultConfig
	_ "github.com/taouniverse/tao-hello"
)

func main() {
	//tao.DevelopMode()		// 2. default configs
	//tao.SetConfigPath("./conf/config_.yml") // 3. custom path

	err := tao.Run(nil, nil)
	if err != nil {
		tao.Error(err)
	}
}
