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
	"time"
)

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
	mushaShugyō := &game{
		id:      "6b91f8d4-42ed-4148-bb20-eb9b31c91eb0",
		name:    "Musha Shugyō",
		created: time.Now(),
		players: []player{player{"Usagi", usagi.id}, player{"Yōjinbō", yōjinbō.id}},
	}
	m.games.id[mushaShugyō.id] = mushaShugyō
	m.games.name[mushaShugyō.name] = mushaShugyō.id
}
