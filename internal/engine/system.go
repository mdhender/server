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

func mksystem(x, y, z int) *System {
	system := &System{
		id:   uuid.New().String(),
		name: fmt.Sprintf("%02d-%02d-%02d", x, y, z),
	}
	system.coords.x = x
	system.coords.y = y
	system.coords.z = z
	return system
}

// System (or Star System) is a group of 1 to 5 Stars.
type System struct {
	id     string
	name   string
	coords struct {
		x int
		y int
		z int
	}
	stars []*Star // a system may have multiple stars
}
