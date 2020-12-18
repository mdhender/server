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

package listing

import (
	"errors"
	"github.com/mdhender/server/internal/auth"
	"time"
)

// Repository defines requirements for fetching data.
type Repository interface {
	GetGame(a *auth.Authorization, id string) (Game, error)
	GetGamePlayer(a *auth.Authorization, id string, name string) (Player, error)
	GetGamePlayers(a *auth.Authorization, id string) ([]Player, error)
	GetGames(a *auth.Authorization, ids ...string) []Game

	GetUser(a *auth.Authorization, id string) (User, error)
	GetUsers(a *auth.Authorization, ids ...string) []User
}

// Service provides listing operations.
type Service interface {
	GetGame(a *auth.Authorization, id string) (Game, error)
	GetGamePlayer(a *auth.Authorization, id string, name string) (Player, error)
	GetGamePlayers(a *auth.Authorization, id string) ([]Player, error)
	GetGames(a *auth.Authorization, ids ...string) []Game

	GetUser(a *auth.Authorization, id string) (User, error)
	GetUsers(a *auth.Authorization, ids ...string) []User
}

// Game defines the properties of a game.
type Game struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Player defines the properties of a player.
type Player struct {
	Name     string `json:"id"`
	UserName string `json:"user_name"`
}

// User defines the properties of a user.
type User struct {
	ID      string    `json:"id"`
	Email   string    `json:"email"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
}

type service struct {
	r Repository
}

// NewService creates a listing service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r: r}
}

// GetGame returns a specific game if the entity is authorized to list that game.
func (s *service) GetGame(a *auth.Authorization, id string) (Game, error) {
	return s.r.GetGame(a, id)
}

// GetGamePlayer returns details for a player in a specific game.
// Returns not found if the entity isn't authorized to list the game or it does not exist.
func (s *service) GetGamePlayer(a *auth.Authorization, id string, name string) (Player, error) {
	return s.r.GetGamePlayer(a, id, name)
}

// GetGamePlayers returns all the players in a specific game that the entity is authorized to list.
// Returns not found if the entity isn't authorized to list the game or it does not exist.
func (s *service) GetGamePlayers(a *auth.Authorization, id string) ([]Player, error) {
	return s.r.GetGamePlayers(a, id)
}

// GetGames returns all games that the entity is authorized to list.
// It never returns an error or a nil list.
func (s *service) GetGames(a *auth.Authorization, ids ...string) []Game {
	return s.r.GetGames(a, ids...)
}

// GetUser returns a specific user if the entity is authorized to list that user.
func (s *service) GetUser(a *auth.Authorization, id string) (User, error) {
	return s.r.GetUser(a, id)
}

// GetUsers returns all users that the entity is authorized to list.
// It never returns an error or a nil list.
func (s *service) GetUsers(a *auth.Authorization, ids ...string) []User {
	return s.r.GetUsers(a, ids...)
}

// ErrGameNotFound is used when the game is not found.
// Note that this could be because the game doesn't exist or the entity making
// the request is not authorized to list the game.
var ErrGameNotFound = errors.New("game not found")

// ErrPlayerNotFound is used when the player is not found.
// Note that this could be because the player doesn't exist or the entity making
// the request is not authorized to list the player.
var ErrPlayerNotFound = errors.New("player not found")

// ErrUserNotFound is used when the user is not found.
// Note that this could be because the user doesn't exist or the entity making
// the request is not authorized to list the user.
var ErrUserNotFound = errors.New("user not found")
