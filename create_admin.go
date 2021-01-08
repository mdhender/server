/*
 * server - a game engine
 * Copyright (C) 2021  Michael D Henderson
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published
 * by the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package engine

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

// CreateAdmin order adds an ID to the state and grants it administrator rights.
type CreateAdmin struct {
	ID string `json:"id"` // id to add with administrator rights
}

// CreateAdmin adds an ID
func (st *State) CreateAdmin(issuedBy, id string) (string, []error) {
	if _, ok := st.admins[issuedBy]; !ok {
		return "", []error{fmt.Errorf("engine refused orders: %w", ERRFORBIDDEN)}
	}
	if id = strings.TrimSpace(id); id == "" {
		id = uuid.New().String()
	}
	if _, ok := st.admins[id]; ok {
		return "", []error{fmt.Errorf("duplicate id: %w", ERRBADREQUEST)}
	}
	st.admins[id] = true
	return id, nil
}
