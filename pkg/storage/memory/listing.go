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
	"github.com/mdhender/server/pkg/auth"
	"github.com/mdhender/server/pkg/listing"
)

// This file implements the listing.Repository interface

// GetAllUsers returns a listing of all users that the call is authorized to list.
// It never returns nil, even if there are no users.
func (m *Store) GetAllUsers(a *auth.Authorization) []listing.User {
	var list []listing.User = []listing.User{}
	isAdmin := a.HasRole("admin")
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

// GetUser returns a listing of a user if the caller is authorized to list that user.
// If the caller is not authorized or the user does not exist, it returns the not found error.
func (m *Store) GetUser(a *auth.Authorization, id string) (listing.User, error) {
	isAuthorized := a.HasRole("admin") || a.ID == id
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
