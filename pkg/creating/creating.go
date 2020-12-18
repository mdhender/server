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

package creating

import (
	"github.com/mdhender/server/pkg/auth"
)

// Repository defines requirements for creating data.
type Repository interface {
	CreateUser(a *auth.Authorization, nu NewUser) (User, error)
}

// Service provides operations to create data.
type Service interface {
	CreateUser(a *auth.Authorization, nu NewUser) (User, error)
}

// NewUser defines the properties of a user to create.
type NewUser struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// User defines the properties of a user.
type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type service struct {
	r Repository
}

// NewService creates a service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r: r}
}

func (s *service) CreateUser(a *auth.Authorization, nu NewUser) (User, error) {
	return s.r.CreateUser(a, nu)
}
