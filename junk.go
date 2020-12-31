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

// Junk order disassembles an asset, reclaiming what it can.
type Junk struct {
	ActorID string `json:"actor_id"` // id of ship or colony being ordered
	AssetID string `json:"asset_id"` // id of ship or colony being junked
}

// Junk disassembles an asset, reclaiming components where possible, recycling where not.
//
// 1. Actor identified by the ActorID must be controlled by the polity issuing the order.
// 2. Asset identified by the AssetID must be controlled by the polity issuing the order.
// 3. Actor and Asset must be in the same Star System.
// 4. Actor and Asset must be within Transport Range of each other.
// 5. The Asset being junked will cease to exist.
func (st *State) Junk(orderedByID, actorID, assetID string) error {
	orderedBy := st.LookupPolity(orderedByID)
	if orderedBy == nil {
		log.Printf("[bug] State.Junk: orderedByID is invalid\n")
		return ERRBUG
	}

	// actor must be a colony or ship controlled by the polity issuing the order
	var actor struct {
		controlledBy *Polity
		colony       *Colony
		ship         *Ship
		system       *System
	}
	if colony := st.LookupColony(actorID); colony != nil {
		actor.controlledBy = colony.controlledBy
		actor.colony = colony
	} else if ship := st.LookupShip(actorID); ship != nil {
		actor.controlledBy = ship.controlledBy
		actor.ship = ship
	} else {
		return fmt.Errorf("invalid actor %q: %w", actorID, ERRBADREQUEST)
	}
	if actor.controlledBy != orderedBy {
		return fmt.Errorf("invalid actor %q: %w", actorID, ERRFORBIDDEN)
	}

	// asset must be a colony or ship controlled by the polity issuing the order
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
	if asset.controlledBy != orderedBy {
		return fmt.Errorf("invalid asset %q: %w", assetID, ERRFORBIDDEN)
	}

	// actor must be in same star system as the asset
	if actor.system != asset.system {
		return fmt.Errorf("actor not in asset's system: %w", ERRFORBIDDEN)
	}

	// actor must be within transport range of the asset
	log.Printf("[todo] State.Junk: assert(actor within transport range of target)\n")

	log.Printf("[bug] State.Junk: not implemented\n")
	return ERRNOTIMPLEMENTED
}
