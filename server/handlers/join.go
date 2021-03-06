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

package handlers

import (
	"os"

	"github.com/boltchat/protocol/events"
	"github.com/boltchat/server/plugins"
	"github.com/boltchat/server/pools"
	"github.com/mitchellh/mapstructure"
)

func HandleJoin(p *pools.ConnPool, c *pools.Connection, e *events.Event) {
	joinData := events.JoinData{}
	mapstructure.Decode(e.Data, &joinData)

	err := plugins.GetManager().HookIdentify(&joinData, c)
	if err != nil {
		c.SendError(err.Error())
		return
	}

	c.User = joinData.User

	motd, hasMotd := os.LookupEnv("MOTD") // Get MOTD env
	if hasMotd {
		// Send MOTD if env var is declared
		c.SendEvent(events.NewMotdEvent(motd))
	}

	p.BroadcastEvent(events.NewJoinEvent(joinData.User))
}
