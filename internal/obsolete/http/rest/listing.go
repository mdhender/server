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
	"github.com/mdhender/server/internal/jsonapi"
	"github.com/mdhender/server/internal/obsolete/auth"
	"github.com/mdhender/server/internal/obsolete/listing"
	"github.com/mdhender/server/internal/way"
	"net/http"
)

// GetGame returns a specific game
func GetGame(ls listing.Service) http.HandlerFunc {
	type okResult struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	a := &auth.Authorization{ID: "usagi", Roles: make(map[string]bool)}
	a.Roles["admin"] = true

	return func(w http.ResponseWriter, r *http.Request) {
		id := way.Param(r.Context(), "id")
		game, err := ls.GetGame(a, id)
		if err != nil {
			if errors.Is(err, listing.ErrGameNotFound) {
				jsonapi.Error(w, r, http.StatusNotFound, err)
				return
			}
			jsonapi.Error(w, r, http.StatusInternalServerError, err)
			return
		}
		jsonapi.Ok(w, r, http.StatusOK, okResult{game.ID, game.Name})
	}
}

// GetGamePlayer returns details for a player in specific game
func GetGamePlayer(ls listing.Service) http.HandlerFunc {
	type okResult struct {
		Name string `json:"name"`
	}

	a := &auth.Authorization{ID: "usagi", Roles: make(map[string]bool)}
	a.Roles["admin"] = true

	return func(w http.ResponseWriter, r *http.Request) {
		id := way.Param(r.Context(), "id")
		name := way.Param(r.Context(), "player_name")
		player, err := ls.GetGamePlayer(a, id, name)
		if err != nil {
			if errors.Is(err, listing.ErrGameNotFound) || errors.Is(err, listing.ErrPlayerNotFound) {
				jsonapi.Error(w, r, http.StatusNotFound, err)
				return
			}
			jsonapi.Error(w, r, http.StatusInternalServerError, err)
			return
		}
		jsonapi.Ok(w, r, http.StatusOK, okResult{player.Name})
	}
}

// GetGamePlayers returns all the players specific game
func GetGamePlayers(ls listing.Service) http.HandlerFunc {
	type okResult []string

	a := &auth.Authorization{ID: "usagi", Roles: make(map[string]bool)}
	a.Roles["admin"] = true

	return func(w http.ResponseWriter, r *http.Request) {
		id := way.Param(r.Context(), "id")
		players, err := ls.GetGamePlayers(a, id)
		if err != nil {
			if errors.Is(err, listing.ErrGameNotFound) {
				jsonapi.Error(w, r, http.StatusNotFound, err)
				return
			}
			jsonapi.Error(w, r, http.StatusInternalServerError, err)
			return
		}
		var list okResult = []string{} // create an empty list since we never return nil
		for _, p := range players {
			list = append(list, p)
		}
		jsonapi.Ok(w, r, http.StatusOK, list)
	}
}

// GetGameSystem returns a specific system in a game
func GetGameSystem(ls listing.Service) http.HandlerFunc {
	type okResult struct {
		Name string `json:"name"`
	}

	a := &auth.Authorization{ID: "usagi", Roles: make(map[string]bool)}
	a.Roles["admin"] = true

	return func(w http.ResponseWriter, r *http.Request) {
		id := way.Param(r.Context(), "id")
		name := way.Param(r.Context(), "system_name")
		system, err := ls.GetGameSystem(a, id, name)
		if err != nil {
			if errors.Is(err, listing.ErrGameNotFound) || errors.Is(err, listing.ErrSystemNotFound) {
				jsonapi.Error(w, r, http.StatusNotFound, err)
				return
			}
			jsonapi.Error(w, r, http.StatusBadRequest, err)
			return
		}
		jsonapi.Ok(w, r, http.StatusOK, okResult{system.Name})
	}
}

// GetGameSystems returns a list of all systems in a game
func GetGameSystems(ls listing.Service) http.HandlerFunc {
	type okResult []string

	a := &auth.Authorization{ID: "usagi", Roles: make(map[string]bool)}
	a.Roles["admin"] = true

	return func(w http.ResponseWriter, r *http.Request) {
		id := way.Param(r.Context(), "id")
		_, err := ls.GetGameSystems(a, id)
		if err != nil {
			if errors.Is(err, listing.ErrGameNotFound) {
				jsonapi.Error(w, r, http.StatusNotFound, err)
				return
			}
			jsonapi.Error(w, r, http.StatusBadRequest, err)
			return
		}
		jsonapi.Ok(w, r, http.StatusOK, okResult{})
	}
}

// GetGames returns all games
func GetGames(ls listing.Service) http.HandlerFunc {
	type detail struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	type okResult []detail

	type formInput struct {
		Data []string `json:"data"`
	}

	a := &auth.Authorization{ID: "usagi", Roles: make(map[string]bool)}
	a.Roles["admin"] = true

	return func(w http.ResponseWriter, r *http.Request) {
		var ids []string
		var list okResult = []detail{} // create an empty list since we never return nil
		for _, game := range ls.GetGames(a, ids...) {
			list = append(list, detail{
				ID:   game.ID,
				Name: game.Name,
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

	a := &auth.Authorization{ID: "usagi", Roles: make(map[string]bool)}
	a.Roles["admin"] = true

	return func(w http.ResponseWriter, r *http.Request) {
		id := way.Param(r.Context(), "id")
		user, err := ls.GetUser(a, id)
		if err != nil {
			if errors.Is(err, listing.ErrUserNotFound) {
				jsonapi.Error(w, r, http.StatusNotFound, err)
				return
			}
			jsonapi.Error(w, r, http.StatusInternalServerError, err)
			return
		}
		jsonapi.Ok(w, r, http.StatusOK, okResult{user.ID, user.Name, user.Email})
	}
}

// GetUsers returns all users
func GetUsers(ls listing.Service) http.HandlerFunc {
	type detail struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	type okResult []detail

	type formInput struct {
		Data []string `json:"data"`
	}

	a := &auth.Authorization{ID: "usagi", Roles: make(map[string]bool)}
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

func GetVersion(ls listing.Service) http.HandlerFunc {
	type okResult struct {
		Version   string `json:"semver"`
		SemVerURL string `json:"semver_url"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		v := ls.GetVersion()

		var versionString string
		if v.PreRelease == "" {
			versionString = fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
		} else if v.Build == "" {
			versionString = fmt.Sprintf("%d.%d.%d-%s", v.Major, v.Minor, v.Patch, v.PreRelease)
		} else {
			versionString = fmt.Sprintf("%d.%d.%d-%s+%s", v.Major, v.Minor, v.Patch, v.PreRelease, v.Build)
		}

		jsonapi.Ok(w, r, http.StatusOK, okResult{
			Version:   versionString,
			SemVerURL: "https://semver.org/",
		})
	}
}
