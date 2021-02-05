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

package identity

import (
	"github.com/boltchat/client/config"
	"github.com/boltchat/client/pgp"
)

// CreateIdentity creates a new Identity.
func CreateIdentity(identity *config.Identity, identityID string) (*Identity, error) {
	identityList := *config.GetIdentityList()
	identityList[identityID] = *identity

	// Create new PGP entity
	entity, createErr := pgp.CreatePGPEntity(identity.Nickname)

	if createErr != nil {
		return nil, createErr
	}

	// Write changes to disk
	_, writeErr := config.IdentityFile.Write(identityList)
	if writeErr != nil {
		return nil, writeErr
	}

	return &Identity{
		Nickname: identity.Nickname,
		Entity:   entity,
	}, nil
}
