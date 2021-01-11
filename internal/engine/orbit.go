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
	"github.com/google/uuid"
)

func mkorbit(star *Star, ring int) *Orbit {
	star.orbits[ring] = &Orbit{
		id:     uuid.New().String(),
		name:   fmt.Sprintf("%s-%02d", star.name, ring),
		system: star.system,
		star:   star,
		ring:   ring,
	}
	return star.orbits[ring]
}

type Orbit struct {
	id       string
	name     string
	system   *System
	star     *Star
	ring     int
	planet   *Planet     // an orbit may have a planet
	deposits []*Resource // or it may have deposits
	colonies []*Colony   // there may be many colonies in the orbit
	ships    []*Ship     // there may be many ships in the orbit
}
