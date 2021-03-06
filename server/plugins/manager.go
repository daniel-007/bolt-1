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

package plugins

import (
	"github.com/boltchat/protocol/events"
	"github.com/boltchat/server/pools"
)

var manager *PluginManager

type PluginManager struct {
	installedPlugins *[]Plugin
}

func (p *PluginManager) Install(plugins ...Plugin) {
	p.installedPlugins = &plugins
}

func (p *PluginManager) GetInstalled() *[]Plugin {
	return p.installedPlugins
}

func (p *PluginManager) HookMessage(msg *events.MessageData, conn *pools.Connection) error {
	for _, plugin := range *p.GetInstalled() {
		err := plugin.OnMessage(msg, conn)

		// Fail fast if a plugin reports an error
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *PluginManager) HookIdentify(data *events.JoinData, conn *pools.Connection) error {
	for _, plugin := range *p.GetInstalled() {
		err := plugin.OnIdentify(data, conn)

		// Fail fast if a plugin reports an error
		if err != nil {
			return err
		}
	}

	return nil
}

func SetManager(mgr *PluginManager) {
	manager = mgr
}

func GetManager() *PluginManager {
	return manager
}
