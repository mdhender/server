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

import "fmt"

// enums
type PlanetKind int

const (
	ASTEROIDBELT PlanetKind = iota
	GASGIANT
	TERRESTRIAL
)

// ColonyKind is TODO
type ColonyKind int

// enums for ColonyKind
const (
	OPEN ColonyKind = iota
	ENCLOSED
	ORBITING
)

// String implements the stringer interface
func (k ColonyKind) String() string {
	switch k {
	case OPEN:
		return "open"
	case ENCLOSED:
		return "enclosed"
	case ORBITING:
		return "orbiting"
	}
	return fmt.Sprint("ColonyKind(%d)", k)
}

// DiplomaticStatus is TODO
type DiplomaticStatus int

// enums for DiplomaticStatus. order is important.
// diplomacy starts at UNKNOWN and increases to ALLY, which is the highest level.
const (
	UNKNOWN      DiplomaticStatus = iota
	ACQUAINTANCE                  // allows messages to be sent
	FRIEND                        // allows assets to be transferred
	ALLY
)

// PopulationKind is the type of population unit.
// It controls what actions the unit may perform.
type PopulationKind int

const (
	OTHERS PopulationKind = iota
	CONSTRUCTION
	PROFESSIONALS
	SOLDIERS
	SPIES
	TRAINEES
	UNSKILLED
)

// String implements the stringer interface
func (k PopulationKind) String() string {
	switch k {
	case CONSTRUCTION:
		return "construction"
	case PROFESSIONALS:
		return "professionals"
	case SOLDIERS:
		return "soldiers"
	case SPIES:
		return "spies"
	case TRAINEES:
		return "trainees"
	case UNSKILLED:
		return "unskilled"
	default: // OTHERS
		return "others"
	}
}

// ResourceKind is TODO
type ResourceKind int

// enums for ResourceKind
const (
	RFUEL ResourceKind = iota
	RGOLD
	RMETAL
	RNONMETAL
)

// String implements the stringer interface
func (k ResourceKind) String() string {
	switch k {
	case RFUEL:
		return "FUEL"
	case RGOLD:
		return "GOLD"
	case RMETAL:
		return "METAL"
	case RNONMETAL:
		return "NONMETAL"
	default:
		return fmt.Sprintf("RESOURCE(%d)", k)
	}
}

// UnitKind is TODO
type UnitKind int

// enums for UnitKind
const (
	NOOP UnitKind = iota
	CONSUMERGOOD
	FARM
	FOOD
	FUEL
	GOLD
	LIGHTSTRUCTURAL
	METAL
	MINE
	NONMETAL
	POPULATION
	POWER
	STRUCTURAL
)

// String implements the stringer interface
func (k UnitKind) String() string {
	switch k {
	case CONSUMERGOOD:
		return "GOODS"
	case FARM:
		return "FARM"
	case FOOD:
		return "FOOD"
	case FUEL:
		return "FUEL"
	case GOLD:
		return "GOLD"
	case LIGHTSTRUCTURAL:
		return "LSU"
	case METAL:
		return "METAL"
	case MINE:
		return "MINE"
	case NONMETAL:
		return "NONMETAL"
	case NOOP:
		return "NOOP"
	case POPULATION:
		return "POP"
	case POWER:
		return "POWER"
	case STRUCTURAL:
		return "SU"
	default:
		return fmt.Sprintf("UNIT(%d)", k)
	}
}
