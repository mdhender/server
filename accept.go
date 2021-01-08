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

package engine

import (
	"fmt"
	"github.com/mdhender/server/internal/polity"
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
// Assumption is that all of the assets of the viceroy were originally controlled by the polity issuing the order.
func (st *State) Accept(issuedByID, assetID string) error {
	if issuedBy := st.LookupPolity(issuedByID); issuedBy == nil {
		log.Printf("[bug] State.Accept: issuedByID is invalid\n")
		return ERRBUG
	}

	// asset must be a colony or ship
	var asset struct {
		controlledBy string
		polity *polity.Polity
		colony       *Colony
		ship         *Ship
	}
	if colony := st.LookupColony(assetID); colony != nil {
		asset.controlledBy = colony.controlledBy.id
		asset.colony = colony
	} else if ship := st.LookupShip(assetID); ship != nil {
		asset.controlledBy = ship.controlledBy.id
		asset.ship = ship
	} else {
		return fmt.Errorf("invalid actor %q: %w", assetID, ERRBADREQUEST)
	}

	// asset must be controlled by a viceroy of the order issuer's polity
	if asset.polity = st.LookupPolity(asset.controlledBy); asset.polity == nil {
		log.Printf("[bug] State.Accept: controlledBy is invalid\n")
		return ERRBUG
	} else if !asset.polity.Serves(issuedByID) {
		return fmt.Errorf("asset refuses order: %w", ERRFORBIDDEN)
	}

	// all checks have passed, so transfer the colony or ship
	if asset.colony != nil {
		return st.transferColony(asset.colony, asset.polity, issuedByID)
	}
	return st.transferShip(asset.ship, asset.polity, issuedByID)
}
