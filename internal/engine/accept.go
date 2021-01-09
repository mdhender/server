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

// Accept order transfers control of a ship or colony from a Viceroy to the Viceroy's Ruler.
type Accept struct {
	AssetID string `json:"asset_id"` // id of ship or colony being ordered
}

// Accept transfers control of an asset from a viceroy back to the original polity.
//
// 1. Asset identified by AssetID must be a ship or a colony.
// 2. Asset must be controlled by a viceroy of the polity issuing the order.
//
// Assumption is that all of the assets of the viceroy were originally
// controlled by the polity issuing the order.
func (st *State) Accept(issuedByID, assetID string) error {
	issuedBy := st.Polity(issuedByID)
	if issuedBy == nil {
		log.Printf("[bug] State.Accept: issuedByID is invalid\n")
		return ERRBUG
	}

	if colony := st.Colony(assetID); colony != nil {
		// asset must be controlled by a viceroy of the polity issuing the order
		if !colony.polity.isViceroyOf(issuedBy) {
			return fmt.Errorf("asset refuses order: %w", ERRFORBIDDEN)
		}
		from, to := colony.polity, issuedBy
		return st.transferColony(colony, from, to)
	}

	if ship := st.Ship(assetID); ship != nil {
		// asset must be controlled by a viceroy of the polity issuing the order
		if !ship.polity.isViceroyOf(issuedBy) {
			return fmt.Errorf("asset refuses order: %w", ERRFORBIDDEN)
		}
		from, to := ship.polity, issuedBy
		return st.transferShip(ship, from, to)
	}

	return fmt.Errorf("invalid assetID %q: %w", assetID, ERRBADREQUEST)
}
