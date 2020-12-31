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
	"log"
)

// Give order transfers control of an asset between polities.
type Give struct {
	AssetID  string `json:"asset_id"`  // id of ship or colony being ordered
	TargetID string `json:"target_id"` // id of the ship, colony, or polity being given to
}

// Give transfers control of an asset between polities.
//
// 1. Asset identified by AssetID must be a ship or a colony.
// 2. The polity issuing the order must control the asset.
// 3. TargetID must be a ship, colony, or polity.
// 4. If the target is a ship or colony, it must be in the same Star System as the asset.
// 5. The target must be controlled by a diplomatic ally of the polity controlling the asset.
// 6. If the source is a Home Colony, the target must the colony's original polity.
func (st *State) Give(orderedByID, assetID, targetID string) error {
	orderedBy := st.LookupPolity(orderedByID)
	if orderedBy == nil {
		log.Printf("[bug] State.Give: orderedByID is invalid\n")
		return ERRBUG
	}

	// asset being given away must be a colony or ship
	var asset struct {
		controlledBy *Polity
		colony       *Colony
		ship         *Ship
		system       *System
	}
	if colony := st.LookupColony(assetID); colony != nil {
		asset.controlledBy = colony.controlledBy
		asset.colony = colony
		asset.system = colony.system
	} else if ship := st.LookupShip(assetID); ship != nil {
		asset.controlledBy = asset.ship.controlledBy
		asset.ship = ship
		asset.system = ship.system
	} else {
		return fmt.Errorf("invalid asset %q: %w", assetID, ERRBADREQUEST)
	}

	// the polity issuing the order must control the asset being given away
	if orderedBy != asset.controlledBy {
		return fmt.Errorf("asset refuses order: %w", ERRFORBIDDEN)
	}

	// target must be a ship, colony, or polity
	var target struct {
		controlledBy *Polity
		colony       *Colony
		ship         *Ship
		system       *System
	}
	if t, ok := st.Lookup(targetID); ok {
		if colony, ok := t.(*Colony); ok {
			target.controlledBy = colony.controlledBy
			target.system = colony.system
		} else if ship, ok := t.(*Ship); ok {
			target.controlledBy = ship.controlledBy
			target.system = ship.system
		} else if polity, ok := t.(*Polity); ok {
			target.controlledBy = polity
		} else {
			return fmt.Errorf("invalid target %q: %w", targetID, ERRBADREQUEST)
		}
	} else {
		return fmt.Errorf("invalid target %q: %w", targetID, ERRBADREQUEST)
	}

	// if the target is a colony
	if target.colony != nil {
		// it must be in same system as the asset.
		if asset.system != target.system {
			return fmt.Errorf("asset not in target's system: %w", ERRFORBIDDEN)
		}
		// a home colony may be given only to the original owning polity.
		if target.colony.isHomeColony() {
			if target.controlledBy != target.colony.originalPolity {
				return fmt.Errorf("home colony may not be given to target %q: %w", targetID, ERRFORBIDDEN)
			}
			panic(ERRNOTIMPLEMENTED)
		}
	}

	// if the target is a ship
	if target.ship != nil {
		// it must be in same system as the asset
		if asset.system != target.system {
			return fmt.Errorf("asset not in target's system: %w", ERRFORBIDDEN)
		}
	}

	// the target receiving the asset must be an ally of the polity controlling the asset
	if !target.controlledBy.isAllied(asset.controlledBy) {
		return fmt.Errorf("target not allied: %w", ERRFORBIDDEN)
	}

	// all checks have passed, so transfer the colony or ship
	if asset.colony != nil {
		return st.transferColony(asset.colony, asset.controlledBy, target.controlledBy)
	}
	return st.transferShip(asset.ship, asset.controlledBy, target.controlledBy)
}
