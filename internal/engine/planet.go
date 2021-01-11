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

func mkplanet(orbit *Orbit, kind PlanetKind) *Planet {
	planet := &Planet{
		id:     uuid.New().String(),
		name:   fmt.Sprintf("%s", orbit.name),
		system: orbit.star.system,
		star:   orbit.star,
		orbit:  orbit,
		kind:   kind,
	}
	planet.orbit.planet = planet
	return planet
}

type Planet struct {
	id           string
	name         string
	system       *System
	star         *Star
	orbit        *Orbit
	kind         PlanetKind
	habitability int // range from 0 to 25
	deposits     []*Resource
	colonies     []*Colony
}
