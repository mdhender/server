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

package colony

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mdhender/server/internal/population"
	"github.com/mdhender/server/internal/text"
	"github.com/mdhender/server/internal/units"
)

// Colony is a unit containing population and resources to manufacture units
type Colony struct {
	id         string
	kind       Kind
	number     string
	createdBy  string
	Batteries  int // power from a plant that expires at the end of the turn
	Name       string
	Note       text.Text // callers may get/set note directly
	OwnedBy    string
	Population population.Population
	Ration     float64 // percent of a full food allotment to be dispersed each turn

	// stockpiles are hoards of food and goods kept by the population.
	// the goal is to keep four turns of both on hand at all times.
	Stockpiles struct {
		Food, Goods struct {
			Want, Have int
		}
	}

	Farms []units.Unit
	Units []units.Unit
}

type Kind int

const (
	OPEN Kind = iota
	ENCLOSED
	ORBITING
)

// New returns an initialized colony.
func New(k Kind, createdBy, number string, p population.Population) Colony {
	c := Colony{
		id:        uuid.New().String(),
		number:    number,
		kind:      k,
		createdBy: createdBy,
	}
	c.Name = c.number
	c.OwnedBy = c.createdBy
	c.Population = p
	c.Ration = 1
	return c
}

func (c Colony) CreatedBy() string {
	return c.createdBy
}

// ID returns the unique identifier for the colony.
// The id will be unique within the game.
func (c Colony) ID() string {
	return c.id
}

// IsSurface returns true if the colony is not an orbiting colony.
func (c Colony) IsSurface() bool {
	return c.kind != ORBITING
}

// Kind returns the type of colony.
func (c Colony) Kind() Kind {
	return c.kind
}

// HullNumber returns the "hull number" of the colony.
// Hull number will be unique within a polity.
func (c Colony) HullNumber() string {
	return c.number
}

// StructureFactor returns the factor used to calculate the amount of
// "volume" that assembled structural units can safely enclose one on
// this colony. The factor depends on the type of colony; it's lower
// for an open colony and highest for a orbiting colony. Since we use
// it as a divisor, that means that structural units on an open colony
// can enclose more units of volume than the same number on an orbiting
// colony.
//
// Note: All population units must be enclosed by a structure.
// Note: Metallic and Non-metallic resources do not needed to be enclosed in
//       surface colonies. (In effect, these two types of resources can be
//       stored in piles on the ground.)
func (c Colony) StructureFactor() int {
	switch c.kind {
	case OPEN:
		return 1
	case ENCLOSED:
		return 5
	case ORBITING:
		return 8
	}
	panic(fmt.Sprintf("assert(kind != %d)", c.kind))
}

// String implements the stringer interface
func (k Kind) String() string {
	switch k {
	case OPEN:
		return "open"
	case ENCLOSED:
		return "enclosed"
	case ORBITING:
		return "orbiting"
	}
	return fmt.Sprintf("kind(%d)", k)
}
