// boltchat
// Copyright (C) 2021  The boltchat Authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package plugins

import (
	"github.com/bolt-chat/protocol/events"
	"github.com/bolt-chat/server/pools"
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

func (p *PluginManager) HookMessage(msg *events.MessageEvent, conn *pools.Connection) error {
	for _, plugin := range *p.GetInstalled() {
		err := plugin.OnMessage(msg, conn)

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