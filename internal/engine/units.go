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
	"github.com/mdhender/server/pkg/utils"
)

// notes on units...
// 18_446_744_073_709_551_615 // maximum unsigned 64 bit integer
//              4_294_967_295 // maximum unsigned 32 bit integer
//                 10_000_000 // people in one population unit
//                 40_000_000 // people fed by one food unit per turn
//              1_000_000_000 // people fed by FARM-1 per turn

type Unit struct {
	Kind      UnitKind
	TechLevel int
	Quantity  int
	Assembled bool
}

// Add returns the sum of two units.
// It panics if they are not the same type and tech level.
func (u Unit) Add(x Unit) Unit {
	if u.Kind == NOOP {
		return x
	} else if x.Kind == NOOP {
		return u
	} else if u.Kind != x.Kind {
		panic("assert(add kind identical)")
	} else if u.TechLevel != x.TechLevel {
		panic("assert(add tech level identical)")
	}
	return Unit{Kind: u.Kind, TechLevel: u.TechLevel, Quantity: u.Quantity + x.Quantity}
}

// Mass returns a number. Must determine what that means.
func (u Unit) Mass() float64 {
	quantity := u.Quantity
	if u.Kind == POPULATION {
		// quantity is scaled for population, so we use the trick
		// of (qty + (n-1)) / n to round results up
		quantity = (quantity + 99) / 100
	}

	techLevel := float64(u.TechLevel)

	var massPerUnit float64
	switch u.Kind {
	case CONSUMERGOOD:
		massPerUnit = 0.6
	case FARM:
		massPerUnit = (2 * techLevel) + 6
	case FOOD:
		massPerUnit = 6
	case FUEL:
		massPerUnit = 1
	case GOLD:
		massPerUnit = 0.6
	case METAL:
		massPerUnit = 1
	case MINE:
		massPerUnit = (2 * techLevel) + 10
	case MINEGROUP:
		massPerUnit = 0
	case NONMETAL:
		massPerUnit = 1
	case NOOP:
		massPerUnit = 0
	case POPULATION:
		massPerUnit = 1
	case POWER:
		massPerUnit = (2 * techLevel) + 10
	default:
		panic(fmt.Sprintf("assert(kind != %d)", u.Kind))
	}
	return float64(quantity) * massPerUnit
}

// Materials returns the amount of metallic and non-metallic units
// needed to build the unit.
func (u Unit) Materials() (metals, nonMetals float64) {
	techLevel := float64(u.TechLevel)
	switch u.Kind {
	case CONSUMERGOOD:
		return 0.2, 0.4
	case FARM:
		return 4 + techLevel, 2 + techLevel
	case FOOD:
		return 0, 0
	case FUEL:
		return 0, 0
	case GOLD:
		return 0, 0
	case METAL:
		return 0, 0
	case MINE:
		return 5 + techLevel, 5 + techLevel
	case MINEGROUP:
		return 0, 0
	case NONMETAL:
		return 0, 0
	case NOOP:
		return 0, 0
	case POPULATION:
		return 0, 0
	case POWER:
		return 5 + techLevel, 5 + techLevel
	}
	panic(fmt.Sprintf("assert(kind != %d)", u.Kind))
}

// Produce returns the units that this item would produce if this
// item could produce items.
// TODO: need to account for required input (population and resources)
func (u Unit) Produce() Unit {
	switch u.Kind {
	case FARM:
		if u.Assembled {
			if u.TechLevel == 1 {
				return Unit{Kind: FOOD, TechLevel: 1, Quantity: 25 * u.Quantity}
			}
			return Unit{Kind: FOOD, TechLevel: 1, Quantity: 5 * u.TechLevel * u.Quantity}
		}
	case POWER:
		if u.Assembled {
			return Unit{Kind: FUEL, TechLevel: 1, Quantity: u.TechLevel * u.Quantity}
		}
	}
	return Unit{}
}

// Sexpr implements the sexpr interface
func (u Unit) Sexpr() string {
	switch u.Kind {
	case CONSUMERGOOD:
		return fmt.Sprintf("(goods %s)", utils.Commas(u.Quantity))
	case FARM:
		return fmt.Sprintf("(farm (tl %d) (qty %s))", u.TechLevel, utils.Commas(u.Quantity))
	case FOOD:
		return fmt.Sprintf("(food %s)", utils.Commas(u.Quantity))
	case FUEL:
		return fmt.Sprintf("(fuel %s)", utils.Commas(u.Quantity))
	case GOLD:
		return fmt.Sprintf("(gold %s)", utils.Commas(u.Quantity))
	case LIGHTSTRUCTURAL:
		return fmt.Sprintf("(lsu (tl %d) (qty %s))", u.TechLevel, utils.Commas(u.Quantity))
	case METAL:
		return fmt.Sprintf("(metal %s)", utils.Commas(u.Quantity))
	case MINE:
		return fmt.Sprintf("(mine (tl %d) (qty %s))", u.TechLevel, utils.Commas(u.Quantity))
	case MINEGROUP:
		return fmt.Sprintf("(mine-group (tl %d) (deposit %s))", u.TechLevel, "?todo?")
	case NONMETAL:
		return fmt.Sprintf("(non-metal %s)", utils.Commas(u.Quantity))
	case NOOP:
		return "(noop)"
	case POPULATION:
		return fmt.Sprintf("(pop %s)", utils.Commas(u.Quantity))
	case POWER:
		return fmt.Sprintf("(power (tl %d) (qty %s))", u.TechLevel, utils.Commas(u.Quantity))
	case STRUCTURAL:
		return fmt.Sprintf("(su (tl %d) (qty %s))", u.TechLevel, utils.Commas(u.Quantity))
	}
	return fmt.Sprintf("(unit (kind %s) (tl %d) (qty %s))", u.Kind.String(), u.TechLevel, utils.Commas(u.Quantity))
}

// Space returns the maximum number of volume units that this unit
// may enclose.
func (u Unit) Space(structureRatio int) int {
	if structureRatio <= 0 || !u.Assembled || !(u.Kind == LIGHTSTRUCTURAL || u.Kind == STRUCTURAL) {
		return 0
	}
	return (u.Quantity * u.TechLevel * u.TechLevel) / structureRatio
}

// String implements the stringer interface
func (u Unit) String() string {
	if u.Kind != POPULATION {
		return fmt.Sprintf("%s-%d", u.Kind, u.TechLevel)
	}
	return u.Kind.String()
}

// Volume returns an abstract number representing the resources
// required to safely enclose the unit for long term storage.
// The amount is determined by the type, tech level, and quantity
// of the unit. Disassembled units typically take up less "space"
// than when they're assembled.
//
// It might help to think of volume as a standard container size. The
// greater the volume, the more "containers" needed to store the unit.
func (u Unit) Volume() float64 {
	quantity := u.Quantity
	if u.Kind == POPULATION {
		// quantity is scaled for population, so we use the trick
		// of (qty + (n-1)) / n to round results up
		quantity = (quantity + 99) / 100
	}

	techLevel := float64(u.TechLevel)

	var containersPerUnit float64
	switch u.Kind {
	case CONSUMERGOOD:
		containersPerUnit = 0.3
	case FARM:
		containersPerUnit = techLevel + 3
		if u.Assembled {
			containersPerUnit *= 2
		}
	case FOOD:
		containersPerUnit = 3
	case FUEL:
		containersPerUnit = 0.5
	case GOLD:
		containersPerUnit = 0.3
	case METAL:
		containersPerUnit = 0.5
	case MINE:
		containersPerUnit = techLevel + 5
		if u.Assembled {
			containersPerUnit *= 2
		}
	case MINEGROUP:
		containersPerUnit = 0
	case NONMETAL:
		containersPerUnit = 0.5
	case NOOP:
		containersPerUnit = 0
	case POPULATION:
		containersPerUnit = 1
	case POWER:
		containersPerUnit = techLevel + 5
		if u.Assembled {
			containersPerUnit *= 2
		}
	default:
		panic(fmt.Sprintf("assert(kind != %d)", u.Kind))
	}
	return float64(quantity) * containersPerUnit
}
