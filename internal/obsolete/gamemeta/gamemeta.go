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

package gamemeta

import (
	"encoding/json"
	"github.com/mdhender/server/internal/obsolete/games"
	"github.com/mdhender/server/internal/obsolete/users"
	"time"
)

type GameMeta struct {
	ID   string
	Name string // public name of the game (not guaranteed to be unique)
	Game *games.Game
	// this keeps confusing me. the player name maps to a specific user.
	// in the game, the player name maps to a specific tribe.
	Players   map[string]*users.User
	CreatedAt time.Time
}

// addPlayer adds a new player to an existing game.
// If the name is already in user or if the user is already a player
// in the game, an error is returned.
func (meta *GameMeta) addPlayer(user *users.User, name string) error {
	if _, ok := meta.Players[name]; ok {
		return ErrDuplicatePlayer
	}
	for _, u := range meta.Players {
		if u.ID == user.ID {
			return ErrDuplicateUser
		}
	}
	meta.Players[name] = user
	return nil
}

func (meta *GameMeta) MarshalJSON() ([]byte, error) {
	data := struct {
		ID        string            `json:"game_id"`
		Name      string            `json:"game_name"`
		CreatedAt time.Time         `json:"created_at"`
		Players   map[string]string `json:"players,omitempty"`
		Game      *games.Game       `json:"data,omitempty"`
	}{
		ID:        meta.ID,
		Name:      meta.Name,
		CreatedAt: meta.CreatedAt,
		Players:   make(map[string]string),
		Game:      meta.Game,
	}
	for name, user := range meta.Players {
		data.Players[name] = user.ID
	}
	return json.Marshal(&data)
}
