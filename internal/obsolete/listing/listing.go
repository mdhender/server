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
	"github.com/mdhender/server/internal/obsolete/auth"
	"time"
)

// Repository defines requirements for fetching data.
type Repository interface {
	GetGame(a *auth.Authorization, id string) (Game, error)
	GetGamePlayer(a *auth.Authorization, id string, name string) (Player, error)
	GetGamePlayers(a *auth.Authorization, id string) (PlayerList, error)
	GetGameSystem(a *auth.Authorization, id string, name string) (SystemDetail, error)
	GetGameSystems(a *auth.Authorization, id string) (SystemList, error)
	GetGames(a *auth.Authorization, ids ...string) []Game

	GetUser(a *auth.Authorization, id string) (User, error)
	GetUsers(a *auth.Authorization, ids ...string) []User
}

// VersionRepository defines requirements for fetching version information.
type VersionRepository interface {
	GetVersion() Version
}

// Service provides listing operations.
type Service interface {
	GetGame(a *auth.Authorization, id string) (Game, error)
	GetGamePlayer(a *auth.Authorization, id string, name string) (Player, error)
	GetGamePlayers(a *auth.Authorization, id string) (PlayerList, error)
	GetGameSystem(a *auth.Authorization, id string, name string) (SystemDetail, error)
	GetGameSystems(a *auth.Authorization, id string) (SystemList, error)
	GetGames(a *auth.Authorization, ids ...string) []Game

	GetUser(a *auth.Authorization, id string) (User, error)
	GetUsers(a *auth.Authorization, ids ...string) []User

	GetVersion() Version
}

// Game defines the properties of a game.
type Game struct {
	ID   string
	Name string
}

// Player defines the properties of a player.
type Player struct {
	Name     string
	UserName string
}

// PlayerList defines a listing of player summaries.
type PlayerList []string

// SystemDetail defines the properties of a system.
type SystemDetail struct {
	Name string
}

// SystemList defines a listing of system summaries
type SystemList struct {
	Name string
}

// User defines the properties of a user.
type User struct {
	ID      string
	Email   string
	Name    string
	Created time.Time
}

// Version defines the properties of a version.
type Version struct {
	Major      int
	Minor      int
	Patch      int
	PreRelease string
	Build      string
}

type service struct {
	r  Repository
	vr VersionRepository
}

// NewService creates a listing service with the necessary dependencies
func NewService(r Repository, vr VersionRepository) Service {
	return &service{r: r, vr: vr}
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
func (s *service) GetGamePlayers(a *auth.Authorization, id string) (PlayerList, error) {
	return s.r.GetGamePlayers(a, id)
}

// GetGameSystem returns detail on a system within a game.
func (s *service) GetGameSystem(a *auth.Authorization, id string, name string) (SystemDetail, error) {
	return s.r.GetGameSystem(a, id, name)
}

// GetGameSystems returns a listing of all systems within a game.
func (s *service) GetGameSystems(a *auth.Authorization, id string) (SystemList, error) {
	return s.r.GetGameSystems(a, id)
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

// GetVersion returns the version of the game engine.
// That may be meaningless?
func (s *service) GetVersion() Version {
	return s.vr.GetVersion()
}

// ErrGameNotFound is used when the game is not found.
// Note that this could be because the game doesn't exist or the entity making
// the request is not authorized to list the game.
var ErrGameNotFound = errors.New("game not found")

// ErrPlayerNotFound is used when the player is not found.
// Note that this could be because the player doesn't exist or the entity making
// the request is not authorized to list the player.
var ErrPlayerNotFound = errors.New("player not found")

// ErrSystemNotFound is used when the system is not found.
// This could be because the system doesn't exist or the entity making the request
// is not authorized to list it.
var ErrSystemNotFound = errors.New("system not found")

// ErrUserNotFound is used when the user is not found.
// Note that this could be because the user doesn't exist or the entity making
// the request is not authorized to list the user.
var ErrUserNotFound = errors.New("user not found")
