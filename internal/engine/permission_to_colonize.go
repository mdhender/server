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
	"log"
)

// PermissionToColonize order enables a ship to establish a new colony on a planet.
// The permission expires at the end of the current turn.
type PermissionToColonize struct {
	PlanetID string `json:"planet_id"` // id of planet (or orbit?) to be colonized
	ShipID   string `json:"ship_id"`   // id of ship being granted permission
}

// PermissionToColonize order enables a ship to establish a new colony on a planet.
//
// 1. Planet identified by PlanetID must be valid
// 2. Planet must contain a colony that is controlled by the polity issuing the order
// 3. Ship identified by ShipID must be valid
// 4. This is a no-op if the ship already has permission to establish a colony on the planet
//
// Note that permission expires at the end of the current turn.
func (st *State) PermissionToColonize(issuedByID, planetID, shipID string) error {
	issuedBy := st.Polity(issuedByID)
	if issuedBy == nil {
		log.Printf("[bug] State.Give: PermissionToColonize is invalid\n")
		return ERRBUG
	}

	planet := st.Planet(planetID)
	if planet == nil {
		return fmt.Errorf("invalid planet %q: %w", planetID, ERRBADREQUEST)
	}

	ship := st.Ship(shipID)
	if ship == nil {
		return fmt.Errorf("invalid ship %q: %w", shipID, ERRBADREQUEST)
	}

	// the polity issuing the order must control at least one colony on the planet
	hasColony := false
	for _, c := range planet.colonies {
		if c.polity == ship.polity {
			// ship doesn't need permission because its polity has
			// already established a colony on the planet!
			return nil
		}

		if c.polity == issuedBy {
			hasColony = true
		}
	}

	if !hasColony {
		return fmt.Errorf("planet refuses orders: %w", ERRFORBIDDEN)
	}

	// all checks passed
	return st.permissionToColonize(planet, ship)
}
