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
	"fmt"
	"github.com/mdhender/server/internal/auth"
	"github.com/mdhender/server/internal/jsonapi"
	"github.com/mdhender/server/pkg/creating"
	"net/http"
)

func CreateUser(cr creating.Service) http.HandlerFunc {
	type okResult struct {
		ID string `json:"id"`
	}
	type formData struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	a := &auth.Authorization{ID: "mdhender", Roles: make(map[string]bool)}
	a.Roles["admin"] = true

	return func(w http.ResponseWriter, r *http.Request) {
		// enforce Content-Type: application/json; charset=utf-8
		if ct := r.Header.Get("Content-Type"); ct != "application/json; charset=utf-8" {
			jsonapi.Error(w, r, http.StatusBadRequest, fmt.Errorf("content-type expected %q: got %q", "application/json; charset=utf-8", ct))
			return
		}

		// enforce a maximum read of 1MB (2^20 bytes) from the request body.
		dec := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<20))
		dec.DisallowUnknownFields() // reject any request with unknown verbs.

		var input formData
		if err := dec.Decode(&input); err != nil {
			jsonapi.Error(w, r, http.StatusBadRequest, fmt.Errorf("bad json %w", err))
			return
		}
		user, err := cr.CreateUser(a, creating.NewUser{
			ID:    input.ID,
			Name:  input.Name,
			Email: input.Email,
		})
		if err != nil {
			if errors.Is(err, creating.ErrUnauthorized) {
				jsonapi.Error(w, r, http.StatusUnauthorized, err)
				return
			}
			jsonapi.Error(w, r, http.StatusBadRequest, err)
			return
		}

		jsonapi.Ok(w, r, http.StatusOK, okResult{user.ID})
	}
}
