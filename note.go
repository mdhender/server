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

// Note order adds a brief message to be displayed on reports for a ship or colony.
// Text must be UTF-8 and is truncated at 200 runes.
type Note struct {
	TargetID string `json:"target_id"` // id of ship or colony being ordered
	Text     string `json:"text"`      // text to be displayed on Owner's report for the ship or colony
}

func (st *State) Note(orderedByID, targetID string, note Text) error {
	orderedBy := st.LookupPolity(orderedByID)
	if orderedBy == nil {
		log.Printf("[bug] State.Note: orderedByID is invalid\n")
		return ERRBUG
	}

	if note = note.TrimSpace(); note.Length() > 200 {
		return fmt.Errorf("invalid text: %w", ERRBADREQUEST)
	}

	// colony must be controlled by the polity issuing the order
	if colony := st.LookupColony(targetID); colony != nil {
		if colony.controlledBy != orderedBy {
			return fmt.Errorf("invalid target %q: %w", targetID, ERRBADREQUEST)
		}
		return st.assignColonyNote(colony, note)
	}

	// ship must be controlled by the polity issuing the order
	if ship := st.LookupShip(targetID); ship != nil {
		if ship.controlledBy != orderedBy {
			return fmt.Errorf("invalid target %q: %w", targetID, ERRBADREQUEST)
		}
		return st.assignShipNote(ship, note)
	}

	// target is not a ship or colony
	return fmt.Errorf("invalid target %q: %w", targetID, ERRBADREQUEST)
}
