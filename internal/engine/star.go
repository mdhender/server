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
	"github.com/google/uuid"
)

func mkstar(system *System) *Star {
	star := &Star{
		id:     uuid.New().String(),
		system: system,
	}
	switch len(system.stars) {
	case 0:
		// just one star (so far), so no suffix on the name
		star.name = system.name
	case 1:
		// two stars, so add the suffix to the first
		system.stars[0].name = system.name + "A"
		star.name = system.name + "B"
	case 2:
		star.name = system.name + "C"
	case 3:
		star.name = system.name + "D"
	case 4:
		star.name = system.name + "E"
	case 5:
		star.name = system.name + "F"
	case 6:
		star.name = system.name + "G"
	case 7:
		star.name = system.name + "H"
	default:
		panic("assert(len(system.stars) < 8)")
	}
	return star
}

// Star is a single star within a system.
type Star struct {
	id     string
	name   string
	system *System
	// star will have 10 orbits; some may be empty
	orbits [10]*Orbit
}
