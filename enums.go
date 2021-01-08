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

// enums
type ColonyKind int
type DiplomaticStatus int
type PlanetKind int
type ResourceKind int

const (
	OPEN ColonyKind = iota
	ENCLOSED
	ORBITING

	// order is important - diplomacy starts at UNKNOWN and increases
	// to ALLY, which is the highest level.
	UNKNOWN      DiplomaticStatus = iota
	ACQUAINTANCE                  // allows messages to be sent
	FRIEND                        // allows assets to be transferred
	ALLY

	ASTEROIDBELT PlanetKind = iota
	GASGIANT
	TERRESTRIAL

	FUEL ResourceKind = iota
	GOLD
	METAL
	NONMETAL
)

func (e ColonyKind) String() string {
	switch e {
	case OPEN:
		return "open"
	case ENCLOSED:
		return "enclosed"
	case ORBITING:
		return "orbiting"
	}
	panic("!")
}
