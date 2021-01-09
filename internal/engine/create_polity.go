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

// CreatePolity order adds a new Polity.
type CreatePolity struct {
	ID   string `json:"id"` // only an admin may provide a default value
	Name string `json:"name"`
}

// CreatePolity creates a new Polity and adds it to the State.
func (st *State) CreatePolity(issuedBy, id, name string) []error {
	if _, ok := st.admins[issuedBy]; !ok {
		return []error{fmt.Errorf("engine refused orders: %w", ERRFORBIDDEN)}
	}
	if id != strings.TrimSpace(id) {
		return []error{fmt.Errorf("invalid characters in id: %w", ERRBADREQUEST)}
	}
	if id == "" {
		id = uuid.New().String()
	}
	if _, ok := st.colonies[id]; ok {
		return []error{fmt.Errorf("duplicate id: %w", ERRBADREQUEST)}
	} else if _, ok := st.planets[id]; ok {
		return []error{fmt.Errorf("duplicate id: %w", ERRBADREQUEST)}
	} else if _, ok := st.polities[id]; ok {
		return []error{fmt.Errorf("duplicate id: %w", ERRBADREQUEST)}
	} else if _, ok := st.ships[id]; ok {
		return []error{fmt.Errorf("duplicate id: %w", ERRBADREQUEST)}
	} else if _, ok := st.stars[id]; ok {
		return []error{fmt.Errorf("duplicate id: %w", ERRBADREQUEST)}
	} else if _, ok := st.systems[id]; ok {
		return []error{fmt.Errorf("duplicate id: %w", ERRBADREQUEST)}
	}
	if name == "" {
		return []error{fmt.Errorf("missing name: %w", ERRBADREQUEST)}
	} else if cleanName := strings.TrimSpace(sanitize(name)); name != cleanName {
		return []error{fmt.Errorf("invalid characters in name: %w", ERRBADREQUEST)}
	} else {
		upperName := strings.ToUpper(name)
		for _, p := range st.polities {
			if name == p.name || strings.ToUpper(p.name) == upperName {
				return []error{fmt.Errorf("duplicate name %q: %w", name, ERRBADREQUEST)}
			}
		}
	}

	st.polities[id] = &Polity{
		id:   id,
		name: name,
	}

	return nil
}
