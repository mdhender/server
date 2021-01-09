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
	"fmt"
	"github.com/mdhender/server/internal/jsonapi"
	"github.com/mdhender/server/internal/obsolete/auth"
	"github.com/mdhender/server/internal/obsolete/reporting"
	"github.com/mdhender/server/internal/way"
	"log"
	"net/http"
)

// GetGamePlayerPrintout
func GetGamePlayerPrintout(rs reporting.Service) http.HandlerFunc {
	type okResult struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	a := &auth.Authorization{ID: "usagi", Roles: make(map[string]bool)}
	a.Roles["admin"] = true

	return func(w http.ResponseWriter, r *http.Request) {
		id := way.Param(r.Context(), "id")
		playerName := way.Param(r.Context(), "player_name")
		turnNumber := way.Param(r.Context(), "turn_number")
		log.Printf("[reporting] not implemented %q %q %q\n", id, playerName, turnNumber)
		jsonapi.Error(w, r, http.StatusNotImplemented, fmt.Errorf("!implmented"))
	}

}
