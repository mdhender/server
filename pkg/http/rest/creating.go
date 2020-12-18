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
	"github.com/mdhender/server/pkg/creating"
	"github.com/mdhender/server/pkg/jsonapi"
	"net/http"
)

func CreateUser(cr creating.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonapi.Error(w, r, http.StatusNotImplemented, fmt.Errorf("not implemented"))
	}
}
