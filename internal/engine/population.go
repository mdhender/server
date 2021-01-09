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

// Population is the number and type of population within a ship or colony.
type Population struct {
	construction  int
	professionals int
	soldiers      int
	spies         int
	trainees      int
	unskilled     int
	others        int
	total         int
}

// FoodNeededPerTurn returns the number of food units needed to
// fully feed a population for one game turn as well as the minimum
// amount needed to prevent starvation.
//
// The full amount is 1.00 food units per population per turn.
// The minimum amount needed to avoid starvation is is 0.25 food units per population per turn.
func (p Population) FoodNeededPerTurn() (min, full int) {
	min = p.total / 4
	if min*4 != p.total {
		min++
	}
	return min, p.total
}
