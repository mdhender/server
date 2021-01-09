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

// Give order transfers control of an asset between polities.
type Give struct {
	AssetID  string `json:"asset_id"`  // id of colony or ship being ordered
	TargetID string `json:"target_id"` // id of the polity, colony, or ship the asset will be given to
}

// Give transfers control of an asset between polities.
//
// 1. Asset identified by AssetID must be a colony or ship.
// 2. The polity issuing the order must control the asset.
// 3. TargetID must be a ship, colony, or polity.
// 4. If the target is a ship or colony, it must be in the same Star System as the asset.
// 5. The target must be controlled by a diplomatic ally of the polity controlling the asset.
// 6. If the source is a Home Colony, the target must the colony's original polity.
func (st *State) Give(issuedByID, assetID, targetID string) error {
	issuedBy := st.Polity(issuedByID)
	if issuedBy == nil {
		log.Printf("[bug] State.Give: issuedByID is invalid\n")
		return ERRBUG
	}

	var asset, target struct {
		polity *Polity
		colony *Colony
		ship   *Ship
		system *System
	}

	// asset must be a colony or ship.
	if asset.colony = st.Colony(assetID); asset.colony != nil {
		asset.polity, asset.system = asset.colony.polity, asset.colony.system
	} else if asset.ship = st.Ship(assetID); asset.ship != nil {
		asset.polity, asset.system = asset.ship.polity, asset.ship.system
	} else {
		return fmt.Errorf("invalid asset %q: %w", assetID, ERRBADREQUEST)
	}

	// the asset must be owned by the polity issuing the order
	if asset.polity != issuedBy {
		return fmt.Errorf("asset refuses order: %w", ERRFORBIDDEN)
	}

	// target must be a polity, colony, or ship.
	// if the target is a colony or ship, set the system.
	// if the target is a colony or ship, set polity to the polity of the target.
	if target.polity = st.Polity(targetID); target.polity == nil {
		if target.colony = st.Colony(targetID); target.colony != nil {
			target.polity, target.system = target.colony.polity, target.colony.system
		} else if target.ship = st.Ship(targetID); target.ship != nil {
			target.polity, target.system = target.ship.polity, target.ship.system
		} else {
			return fmt.Errorf("invalid target %q: %w", targetID, ERRBADREQUEST)
		}
	}

	// the target must belong to an ally
	if !issuedBy.isAlliedTo(target.polity) {
		return fmt.Errorf("asset refuses order: %w", ERRFORBIDDEN)
	}

	// if the target is a colony or ship, it must be in the same system as the asset.
	if target.system != nil && asset.system != target.system {
		return fmt.Errorf("asset not in target's system: %w", ERRFORBIDDEN)
	}

	// if the asset is a home colony, the target must be the colony's original polity
	if asset.colony.isHomeColony() && target.polity != asset.colony.originalPolity {
		return fmt.Errorf("asset refuses order: %w", ERRFORBIDDEN)
	}

	// all checks have passed, so transfer the colony or ship
	if asset.colony == nil {
		return st.transferShip(asset.ship, asset.polity, target.polity)
	}
	return st.transferColony(asset.colony, asset.polity, target.polity)
}
