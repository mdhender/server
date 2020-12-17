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

import (
	"time"
)

func New() (*Store, error) {
	m := &Store{}
	m.users.id = make(map[string]*user)
	m.users.name = make(map[string]string)
	return m, nil
}

// MockData based on Stan Sakai's classic Usagi Yojimbo.
//   https://stansakai.com/
//   http://www.usagiyojimbo.com/
func (m *Store) MockData() {
	usagi := &user{
		id:      "bf4c8168-6aab-409d-80cf-a4ee901904ef",
		email:   "usagi@server.example.com",
		name:    "usagi",
		roles:   []string{"admin", "user"},
		created: time.Now(),
	}
	m.users.id[usagi.id] = usagi
	m.users.name[usagi.name] = usagi.id

	yōjinbō := &user{
		id:      "236bb1a5-1ae8-411a-a71f-791f4f03aa99",
		email:   "yōjinbō@server.example.com",
		name:    "yōjinbō",
		roles:   []string{"user"},
		created: time.Now(),
	}
	m.users.id[yōjinbō.id] = yōjinbō
	m.users.name[yōjinbō.name] = yōjinbō.id

	// game named Musha Shugyō
}

type Store struct {
	games struct {
		// id is a map from game id to game properties
		id map[string]*game
		// name is a map from game name to game id
		name map[string]string
	}
	users struct {
		// id is a map from user id to user properties
		id map[string]*user
		// name is a map from user name to user id
		name map[string]string
	}
}
