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

func New() (*Store, error) {
	m := &Store{}
	m.users.data = make(map[string]*user)
	return m, nil
}

type Store struct {
	users struct {
		// data is a map from user name to user properties
		data map[string]*user
	}
}
