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
	"encoding/json"
	"errors"
	"github.com/mdhender/server/pkg/auth"
	"github.com/mdhender/server/pkg/jsonapi"
	"github.com/mdhender/server/pkg/listing"
	"github.com/mdhender/server/pkg/way"
	"net/http"
)

// GetUser returns a specific user
func GetUser(ls listing.UserService) http.HandlerFunc {
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

// GetUsers returns all users
func GetUsers(ls listing.UserService) http.HandlerFunc {
	type detail struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	type okResult []detail

	type formInput struct {
		Data []string `json:"data"`
	}

	a := &auth.Authorization{ID: "mdhender", Roles: make(map[string]bool)}
	a.Roles["admin"] = true

	return func(w http.ResponseWriter, r *http.Request) {
		var ids []string
		if r.Method == "POST" { // support sending a list of ids to fetch
			// Enforce a maximum read of 1MB from the request body.
			dec := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<20))
			// reject any request with unknown verbs.
			dec.DisallowUnknownFields()
			var input formInput
			if err := dec.Decode(&formInput{}); err != nil {
				jsonapi.Error(w, r, http.StatusBadRequest, err)
				return
			}
			if len(input.Data) != 0 {
				ids = input.Data
			}
		}

		var list okResult = []detail{} // create an empty list since we never return nil
		for _, user := range ls.GetUsers(a, ids...) {
			list = append(list, detail{
				ID:    user.ID,
				Email: user.Email,
				Name:  user.Name,
			})
		}
		jsonapi.Ok(w, r, http.StatusOK, list)
	}
}
