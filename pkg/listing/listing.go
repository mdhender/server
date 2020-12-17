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
	"github.com/mdhender/server/pkg/auth"
	"time"
)

// UserRepository defines requirements for fetching data.
type Repository interface {
	GetAllUsers(a *auth.Authorization) []User
	GetUser(a *auth.Authorization, id string) (User, error)
}

// Service provides listing operations.
type Service interface {
	GetAllUsers(a *auth.Authorization) []User
	GetUser(a *auth.Authorization, id string) (User, error)
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
	return &service{r}
}

// GetAllUsers returns all users that the entity is authorized to list.
// It never returns an error or a nil list.
func (s *service) GetAllUsers(a *auth.Authorization) []User {
	return s.r.GetAllUsers(a)
}

// GetUser returns a specific user if the entity is authorized to list that user.
func (s *service) GetUser(a *auth.Authorization, id string) (User, error) {
	return s.r.GetUser(a, id)
}

// ErrUserNotFound is used when the user is not found.
// Note that this could be because the user doesn't exist or the entity making
// the request is not authorized to list the user.
var ErrUserNotFound = errors.New("user not found")
