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

// CreateCluster order creates a new Cluster.
type CreateCluster struct {
	ID       string `json:"id,omitempty"` // only an admin may provide a default value
	Name     string `json:"name,omitempty"`
	Polities []struct {
		ID       string  `json:"id,omitempty"`
		Name     string  `json:"name"`
		Scarcity float64 `json:"scarcity,omitempty"` // percentage 0.25 to 1.00
	} `json:"polities"`
}

type clusterPolities struct {
	id       string
	name     string
	scarcity float64
}

// CreateCluster creates a new Cluster.
func (st *State) CreateCluster(issuedBy, id, name string, polities ...clusterPolities) []error {
	if _, ok := st.admins[issuedBy]; !ok {
		return []error{fmt.Errorf("engine refused orders: %w", ERRFORBIDDEN)}
	}

	if id != strings.TrimSpace(id) {
		return []error{fmt.Errorf("invalid characters in id: %w", ERRBADREQUEST)}
	}
	if id == "" {
		id = uuid.New().String()
	}

	for i := range polities {
		p := polities[i]
		if p.id == "" || p.id != strings.TrimSpace(sanitize(p.id)) {
			p.id = uuid.New().String()
		}
		if p.name == "" || p.name != strings.TrimSpace(sanitize(p.name)) {
			p.name = fmt.Sprintf("POLITY-%02d", i+1)
		}
		if p.scarcity < 0.25 {
			p.scarcity = 0.25
		} else if p.scarcity > 1 {
			p.scarcity = 1
		}
		polities[i] = p
	}

	cluster := mkcluster()
	st.admins = cluster.admins
	st.polities = cluster.polities
	st.systems = cluster.systems
	st.stars = cluster.stars
	st.planets = cluster.planets
	st.colonies = cluster.colonies
	st.ships = cluster.ships

	return nil
}
