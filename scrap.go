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

// Scrap order...
type Scrap struct {
	ActorID   string `json:"actor_id"` // id of ship or colony being ordered
	Item      string `json:"item"`
	TechLevel int    `json:"tech_level"`
	Quantity  int    `json:"quantity"`
}

// Scrap disassembles an item, recycling components into resources.
//
// 1. Actor identified by the ActorID must be controlled by the polity issuing the order.
// 2. The Item to scrap must be Unassembled or a Non-Assembly unit.
// 3. Quantity must not be less than zero (0).
// 4. Quantity may exceed the actual number of items controlled by the actor.
//    Any overage will be ignored.
// 5. One (1) Constructor is required per 300 Mass Units (or portion).
// 6. The Item being scrapped will lose 30% of its Mass as waste.
func (st *State) Scrap(orderedByID, actorID, item string, techLevel, quantity int) error {
	orderedBy := st.LookupPolity(orderedByID)
	if orderedBy == nil {
		log.Printf("[bug] State.Scrap: orderedByID is invalid\n")
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

	log.Printf("[bug] State.Scrap: not implemented\n")
	return ERRNOTIMPLEMENTED
}
