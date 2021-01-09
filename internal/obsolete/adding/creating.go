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

package adding

import (
	"errors"
	"github.com/mdhender/server/internal/obsolete/auth"
)

// Repository defines requirements for creating data.
type Repository interface {
	AddGame(a *auth.Authorization, ng NewGame) (Game, error)
	AddUser(a *auth.Authorization, nu NewUser) (User, error)
}

// Service provides operations to create data.
type Service interface {
	AddGame(a *auth.Authorization, ng NewGame) (Game, error)
	AddUser(a *auth.Authorization, nu NewUser) (User, error)
}

// Game defines the properties of a game.
type Game struct {
	ID   string
	Name string
}

// NewGame defines the properties of a game to create.
type NewGame struct {
	ID   string
	Name string
}

// NewUser defines the properties of a user to create.
type NewUser struct {
	ID    string
	Email string
	Name  string
}

// User defines the properties of a user.
type User struct {
	ID    string
	Email string
	Name  string
}

type service struct {
	r Repository
}

// NewService creates a service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r: r}
}

func (s *service) AddGame(a *auth.Authorization, ng NewGame) (Game, error) {
	return s.r.AddGame(a, ng)
}

func (s *service) AddUser(a *auth.Authorization, nu NewUser) (User, error) {
	return s.r.AddUser(a, nu)
}

// ErrDuplicateEmail is used when the e-mail address is not unique.
var ErrDuplicateEmail = errors.New("duplicate e-mail address")

// ErrDuplicateID is used when the id is not unique.
var ErrDuplicateID = errors.New("duplicate id")

// ErrDuplicateName is used when the user name is not unique.
var ErrDuplicateName = errors.New("duplicate user name")

// ErrInvalidEmail is used when the email is not valid.
var ErrInvalidEmail = errors.New("invalid e-mail address")

// ErrInvalidID is used when the id is not valid.
var ErrInvalidID = errors.New("invalid id")

// ErrInvalidName is used when the name is not valid.
var ErrInvalidName = errors.New("invalid name")

// ErrUnauthorized is used when the caller does not have the role needed for an action.
var ErrUnauthorized = errors.New("unauthorized")
