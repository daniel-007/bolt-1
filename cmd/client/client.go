// Copyright 2021 The boltchat Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"

	"github.com/boltchat/client/cli"
	"github.com/boltchat/client/config"
)

func main() {
	// Load the config
	config.LoadConfig()
	config.LoadIdentityList()

	cmd, cmdErr := cli.ParseCommand(os.Args[1:])
	if cmdErr != nil {
		cli.HandleCommandError(cmdErr)
	}

	execErr := cmd.Execute()
	if execErr != nil {
		cli.HandleCommandError(execErr)
	}
}
