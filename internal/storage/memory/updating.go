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
	"fmt"
	"github.com/mdhender/server/internal/obsolete/auth"
	"github.com/mdhender/server/internal/obsolete/updating"
)

// This file implements the updating.Repository interface

// UpdateGame applies changes to an existing game to the store.
func (m *Store) UpdateGame(a *auth.Authorization, gu updating.GameUpdates) error {
	isAdmin := a.HasRole("admin")
	if !isAdmin {
		return updating.ErrNotAuthorized
	}

	return fmt.Errorf("not implemented")
}

// UpdateGameOrders applies a new set of orders to an existing game to the store.
func (m *Store) UpdateGameOrders(a *auth.Authorization, o updating.Orders) error {
	isAdmin := a.HasRole("admin")
	if !isAdmin {
		return updating.ErrNotAuthorized
	}

	return fmt.Errorf("not implemented")
}
