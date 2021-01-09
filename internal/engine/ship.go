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

import "github.com/mdhender/server/internal/obsolete/population"

type Ship struct {
	id         string
	polity     *Polity
	number     string
	system     *System
	homePort   *Colony
	name       string
	note       Text
	population population.Population
	units      struct {
		farms []FarmUnit
	}
	// percent of a full food allotment to be dispersed each turn
	ration float64
}
