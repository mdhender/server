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
	"unicode/utf8"
)

// Name order changes the name assigned to a ship, colony, star, system, planet, or polity.
// The "Type" parameter helps prevent accidental renames. The value must match the type of
// the entity being renamed or the order will be refused.
type Name struct {
	EntityID string `json:"entity_id"` // id of entity being ordered
	Type     string `json:"type"`      // type of entity to assign name to
	Name     string `json:"name"`      // name to assign to the entity
}

// Name changes the name assigned to a ship, colony, star, system, planet, or polity.
//
// 1. Entity identified by EntityID
// 2. The name will be converted to valid UTF-8
// 3. Certain prohibited characters in the name will be replaced with underscores ('_')
// 4. The name must not be empty.
// 5. The name must not be longer than fifty (50) characters after conversion and cleaning.
// 6. Type must match exactly the type of entity being named.
// 7. If the entity is a ship, colony, or polity, it must be controlled by the issuer of the order.
// 8. Names of ships and colonies are not required to be unique.
//
// Note: each Polity maintains its own "database" of unique names for Stars, Star Systems, and Planets.
func (st *State) Name(orderedByID, entityID, typeFlag, name string) error {
	orderedBy := st.LookupPolity(orderedByID)
	if orderedBy == nil {
		log.Printf("[bug] State.Name: orderedByID is invalid\n")
		return ERRBUG
	}

	entity, ok := st.Lookup(entityID)
	if !ok {
		return fmt.Errorf("invalid entity %q: %w", entityID, ERRBADREQUEST)
	}

	name = sanitize(name)
	if n := utf8.RuneCountInString(name); n == 0 || n > 50 {
		return fmt.Errorf("invalid name: %w", ERRBADREQUEST)
	}

	if colony, ok := entity.(*Colony); ok {
		if typeFlag != "colony" {
			return fmt.Errorf("invalid type %q: %w", typeFlag, ERRBADREQUEST)
		}
		if colony.controlledBy != orderedBy {
			return fmt.Errorf("colony refuses order: %w", ERRFORBIDDEN)
		}
		return st.assignColonyName(colony, name)
	} else if planet, ok := entity.(*Planet); ok {
		if typeFlag != "planet" {
			return fmt.Errorf("invalid type %q: %w", typeFlag, ERRBADREQUEST)
		}
		return st.assignPlanetName(planet, name)
	} else if polity, ok := entity.(*Polity); ok {
		if typeFlag != "polity" {
			return fmt.Errorf("invalid type %q: %w", typeFlag, ERRBADREQUEST)
		}
		if polity != orderedBy {
			return fmt.Errorf("polity refuses order: %w", ERRFORBIDDEN)
		}
		return st.assignPolityName(polity, name)
	} else if ship, ok := entity.(*Ship); ok {
		if typeFlag != "ship" {
			return fmt.Errorf("invalid type %q: %w", typeFlag, ERRBADREQUEST)
		}
		if ship.controlledBy != orderedBy {
			return fmt.Errorf("ship refuses order: %w", ERRFORBIDDEN)
		}
		return st.assignShipName(ship, name)
	} else if star, ok := entity.(*Star); ok {
		if typeFlag != "star" {
			return fmt.Errorf("invalid type %q: %w", typeFlag, ERRBADREQUEST)
		}
		return st.assignStarName(star, name)
	} else if system, ok := entity.(*System); ok {
		if typeFlag != "system" {
			return fmt.Errorf("invalid type %q: %w", typeFlag, ERRBADREQUEST)
		}
		return st.assignSystemName(system, name)
	}

	// if we fall through to here, it can only be because we weren't given a valid entity
	return fmt.Errorf("entity can not be named: %w", ERRBADREQUEST)
}
