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

// game defines the properties of a game.
type game struct {
	id        string
	name      string
	created   time.Time
	completed time.Time
	players   []player
}

// player defines the properties of a player,
// which is an instance of a user in a game.
type player struct {
	name string // name of player
	user string // name of user
	// todo: more information on the player. stuff like race, possessions, etc
}

// user defines the properties of a user.
type user struct {
	id      string
	email   string
	name    string
	roles   []string
	created time.Time
}
