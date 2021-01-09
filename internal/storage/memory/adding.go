// server - a game engine
// Copyright (C) 2020  Michael D Henderson
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package memory

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mdhender/server/internal/obsolete/adding"
	"github.com/mdhender/server/internal/obsolete/auth"
	"strings"
)

// This file implements the adding.Repository interface

// AddGame adds a new game to the store.
// If the caller is an admin, then the request may specify the ID to use.
func (m *Store) AddGame(a *auth.Authorization, ng adding.NewGame) (adding.Game, error) {
	isAdmin := a.HasRole("admin")
	if !isAdmin {
		return adding.Game{}, adding.ErrUnauthorized
	}

	return adding.Game{}, fmt.Errorf("not implemented")
}

// AddUser adds a new user to the store.
// If the caller is an admin, then the request may specify the ID to use.
// TODO: Only admins can call this, so why do the test for ID first?
func (m *Store) AddUser(a *auth.Authorization, nu adding.NewUser) (adding.User, error) {
	isAdmin := a.HasRole("admin")
	if !isAdmin {
		return adding.User{}, adding.ErrUnauthorized
	}

	if nu.Name == "" || strings.TrimSpace(nu.Name) != nu.Name {
		return adding.User{}, adding.ErrInvalidName
	} else if nu.Email == "" || strings.TrimSpace(nu.Email) != nu.Email {
		return adding.User{}, adding.ErrInvalidEmail
	}

	m.users.Lock()
	defer m.users.Unlock()

	var id string
	if nu.ID == "" {
		id = uuid.New().String()
	} else if isAdmin {
		if strings.TrimSpace(nu.ID) != nu.ID {
			return adding.User{}, adding.ErrInvalidID
		}
		id = nu.ID
	}

	// confirm that we don't duplicate any keys
	if _, ok := m.users.id[id]; ok {
		return adding.User{}, adding.ErrDuplicateID
	} else if _, ok := m.users.name[nu.Name]; ok {
		return adding.User{}, adding.ErrDuplicateName
	} else if _, ok := m.users.email[nu.Email]; ok {
		return adding.User{}, adding.ErrDuplicateEmail
	}

	user := &user{
		id:    id,
		email: nu.Email,
		name:  nu.Name,
	}

	m.users.id[user.id] = user
	m.users.email[user.email] = user.id
	m.users.name[user.name] = user.id

	return adding.User{
		ID:    id,
		Email: nu.Email,
		Name:  nu.Name,
	}, nil
}
