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
	"github.com/mdhender/server/internal/auth"
	"github.com/mdhender/server/internal/listing"
)

// This file implements the listing.Repository interface

// GetGame returns a listing of a game if the caller is authorized to list that game.
// If the caller is not authorized or the game does not exist, it returns the not found error.
func (m *Store) GetGame(a *auth.Authorization, id string) (listing.Game, error) {
	isAdmin := a.HasRole("admin")
	if isAdmin {
		if game, ok := m.games.id[id]; ok {
			return listing.Game{
				ID:   game.id,
				Name: game.name,
			}, nil
		}
	}
	return listing.Game{}, listing.ErrGameNotFound
}

// GetGamePlayer returns data for a player in a game.
// If the caller is not authorized or the game/player does not exist, it returns the not found error.
func (m *Store) GetGamePlayer(a *auth.Authorization, id, name string) (listing.Player, error) {
	isAdmin := a.HasRole("admin")
	if isAdmin {
		if game, ok := m.games.id[id]; ok {
			for _, player := range game.players {
				if player.name == name {
					return listing.Player{
						Name:     player.name,
						UserName: player.user,
					}, nil
				}
			}
			return listing.Player{}, listing.ErrPlayerNotFound
		}
	}
	return listing.Player{}, listing.ErrGameNotFound
}

// GetGamePlayers returns data for all players in a game.
// If the caller is not authorized or the game/player does not exist, it returns the not found error.
func (m *Store) GetGamePlayers(a *auth.Authorization, id string) ([]listing.Player, error) {
	var list []listing.Player = []listing.Player{}
	isAdmin := a.HasRole("admin")
	if isAdmin {
		if game, ok := m.games.id[id]; ok {
			for _, player := range game.players {
				list = append(list, listing.Player{
					Name:     player.name,
					UserName: player.user,
				})
			}
			return list, nil
		}
	}
	return list, listing.ErrGameNotFound
}

// GetGames returns a listing of games that the call is authorized to list.
// If the list of ids passed in is empty, we attempt to return all games.
// Otherwise, we return only the games in the list.
// We never return nil, even if there are no games.
func (m *Store) GetGames(a *auth.Authorization, ids ...string) []listing.Game {
	var list []listing.Game = []listing.Game{}
	isAdmin := a.HasRole("admin")
	if len(ids) == 0 { // this is a request for all games
		for _, game := range m.games.id {
			isAuthorized := isAdmin || a.ID == game.id
			if isAuthorized {
				list = append(list, listing.Game{
					ID:   game.id,
					Name: game.name,
				})
			}
		}
		return list
	}
	// a request for a specific set of games
	for _, id := range ids {
		isAuthorized := isAdmin || a.ID == id
		if isAuthorized {
			if game, ok := m.games.id[id]; ok {
				list = append(list, listing.Game{
					ID:   game.id,
					Name: game.name,
				})
			}
		}
	}
	return list
}

// GetUser returns a listing of a user if the caller is authorized to list that user.
// If the caller is not authorized or the user does not exist, it returns the not found error.
func (m *Store) GetUser(a *auth.Authorization, id string) (listing.User, error) {
	isAdmin := a.HasRole("admin")
	isAuthorized := isAdmin || a.ID == id
	if isAuthorized {
		if user, ok := m.users.id[id]; ok {
			return listing.User{
				ID:      user.id,
				Email:   user.email,
				Name:    user.name,
				Created: user.created,
			}, nil
		}
	}
	return listing.User{}, listing.ErrUserNotFound
}

// GetUsers returns a listing of users that the call is authorized to list.
// If the list of ids passed in is empty, we attempt to return all users.
// Otherwise, we return only the users in the list.
// We never return nil, even if there are no users.
func (m *Store) GetUsers(a *auth.Authorization, ids ...string) []listing.User {
	var list []listing.User = []listing.User{}
	isAdmin := a.HasRole("admin")
	if len(ids) == 0 { // this is a request for all users
		for _, user := range m.users.id {
			isAuthorized := isAdmin || a.ID == user.id
			if isAuthorized {
				list = append(list, listing.User{
					ID:      user.id,
					Email:   user.email,
					Name:    user.name,
					Created: user.created,
				})
			}
		}
		return list
	}
	// a request for a specific set of users
	for _, id := range ids {
		isAuthorized := isAdmin || a.ID == id
		if isAuthorized {
			if user, ok := m.users.id[id]; ok {
				list = append(list, listing.User{
					ID:      user.id,
					Email:   user.email,
					Name:    user.name,
					Created: user.created,
				})
			}
		}
	}
	return list
}
