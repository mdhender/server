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

// HomePortChange order reassigns the home port of a ship
type HomePortChange struct {
	ShipID   string `json:"ship_id"`   // id of ship being ordered
	ColonyID string `json:"colony_id"` // id of colony being targeted
}

// HomePortChange assigns the home port of a ship to a colony.
//
// 1. Ship identified by the ShipID must be controlled by the polity issuing the order.
// 2. Colony identified by the ColonyID must be controlled by the polity issuing the order.
func (st *State) HomePortChange(issuedByID, shipID, colonyID string) error {
	issuedBy := st.Polity(issuedByID)
	if issuedBy == nil {
		log.Printf("[bug] State.HomePortChange: issuedByID is invalid\n")
		return ERRBUG
	}

	// ship must be a ship controlled by the polity issuing the order
	ship := st.Ship(shipID)
	if ship == nil {
		return fmt.Errorf("invalid ship %q: %w", shipID, ERRBADREQUEST)
	} else if ship.polity != issuedBy {
		return fmt.Errorf("ship refuses order: %w", ERRFORBIDDEN)
	}

	// colony must be a colony controlled by the polity issuing the order
	colony := st.Colony(colonyID)
	if colony == nil {
		return fmt.Errorf("invalid colony %q: %w", colonyID, ERRBADREQUEST)
	} else if colony.polity != issuedBy {
		return fmt.Errorf("colony refuses order: %w", ERRFORBIDDEN)
	}

	// all checks have passed, so assign the ship to the colony
	return st.setHomePort(ship, colony)
}
