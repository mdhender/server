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

package updating

import (
	"errors"
	"github.com/mdhender/server/internal/auth"
)

type Repository interface {
	UpdateGame(a *auth.Authorization, g GameUpdates) error
	UpdateGameOrders(a *auth.Authorization, o Orders) error
}

type Service interface {
	UpdateGame(a *auth.Authorization, g GameUpdates) error
	UpdateGameOrders(a *auth.Authorization, o Orders) error
}

type GameUpdates struct {
	ID string
}

type Orders struct {
	ID string
}

func NewService(r Repository) Service {
	return &service{r: r}
}

type service struct {
	r Repository
}

func (s *service) UpdateGame(a *auth.Authorization, g GameUpdates) error {
	return s.r.UpdateGame(a, g)
}

func (s *service) UpdateGameOrders(a *auth.Authorization, o Orders) error {
	return s.r.UpdateGameOrders(a, o)
}

// ErrNotAuthorized is used when the entity making
// the request is not authorized to update the game.
var ErrNotAuthorized = errors.New("not authorized")
