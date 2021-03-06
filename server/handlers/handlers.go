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
	"github.com/boltchat/protocol/events"
	"github.com/boltchat/server/pools"
)

type handler = func(p *pools.ConnPool, c *pools.Connection, e *events.Event)

var handlerMap = map[events.Type]handler{
	events.MessageType: HandleMessage,
	events.JoinType:    HandleJoin,
	events.CommandType: HandleCommand,
}

func GetHandler(evtType events.Type) handler {
	if evtHandler, ok := handlerMap[evtType]; ok {
		return evtHandler
	}

	// Use default handler if event is not recognized
	return HandleDefault
}
