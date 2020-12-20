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

// Package memory implements in-memory data storage.
package memory

import "sync"

func New() (*Store, error) {
	m := &Store{}

	m.games.id = make(map[string]*game)
	m.games.name = make(map[string]string)
	m.users.id = make(map[string]*user)
	m.users.email = make(map[string]string)
	m.users.name = make(map[string]string)
	m.version = version{0,0,1,"",""}

	return m, nil
}

type Store struct {
	games struct {
		// id is a map from game id to game properties
		id map[string]*game
		// name is a map from game name to game id
		name map[string]string
	}
	users struct {
		sync.RWMutex
		// id is a map from user id to user properties
		id map[string]*user
		// email is a map from user email to user id
		email map[string]string
		// name is a map from user name to user id
		name map[string]string
	}
	version version
}
