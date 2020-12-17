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

package rest

import (
	"errors"
	"github.com/mdhender/server/pkg/auth"
	"github.com/mdhender/server/pkg/jsonapi"
	"github.com/mdhender/server/pkg/listing"
	"github.com/mdhender/server/pkg/way"
	"net/http"
)

// GetAllUsers returns all users
func GetAllUsers(ls listing.Service) http.HandlerFunc {
	type detail struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	type okResult []detail

	a := &auth.Authorization{ID: "mdhender", Roles: make(map[string]bool)}
	a.Roles["admin"] = true

	return func(w http.ResponseWriter, r *http.Request) {
		var list okResult = []detail{} // create an empty list since we never return nil
		for _, user := range ls.GetAllUsers(a) {
			list = append(list, detail{
				ID:    user.ID,
				Email: user.Email,
				Name:  user.Name,
			})
		}
		jsonapi.Ok(w, r, http.StatusOK, list)
	}
}

// GetUser returns a specific user
func GetUser(ls listing.Service) http.HandlerFunc {
	type okResult struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	a := &auth.Authorization{ID: "mdhender", Roles: make(map[string]bool)}
	a.Roles["admin"] = true

	return func(w http.ResponseWriter, r *http.Request) {
		id := way.Param(r.Context(), "id")
		user, err := ls.GetUser(a, id)
		if err != nil {
			if errors.Is(err, listing.ErrUserNotFound) {
				jsonapi.Error(w, r, http.StatusNotFound, listing.ErrUserNotFound)
				return
			}
			jsonapi.Error(w, r, http.StatusInternalServerError, err)
			return
		}
		jsonapi.Ok(w, r, http.StatusOK, okResult{user.ID, user.Name, user.Email})
	}
}
