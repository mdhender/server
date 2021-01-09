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

type Colony struct {
	id                string
	kind              ColonyKind
	number            string
	polity            *Polity
	originalPolity    *Polity // set only if this is a Home Colony
	system            *System
	name              string
	note              Text
	population        Population
	foodStockpileGoal int
	units             []Unit
	rebels            struct {
		construction  float64
		professionals float64
		soldiers      float64
		spies         float64
		trainees      float64
		unskilled     float64
		others        float64
	}
	storage struct {
		food     int
		fuel     int
		gold     int
		metal    int
		nonmetal int
	}
	// percent of a full food allotment to be dispersed each turn
	ration   float64
	controls struct {
		ships map[string]*Ship // acts as home port to
	}
	batteries int // power from a plant that expires at the end of the turn
}

func (c *Colony) addShip(s *Ship) {
	if c == nil || s == nil {
		return
	}
	c.controls.ships[s.id] = s
	s.homePort = c
}

func (c *Colony) delShip(s *Ship) {
	if c == nil || s == nil {
		return
	}
	delete(c.controls.ships, s.id)
	// todo: this should be set to the polity's home world
	s.homePort = nil
}

func (c *Colony) isHomeColony() bool {
	if c == nil {
		return false
	}
	return c.originalPolity != nil
}

func (c *Colony) setPolity(p *Polity) {
	c.polity = p
	p.addColony(c)
}

func (c *Colony) xferShip(s *Ship) error {
	s.homePort.delShip(s)
	c.addShip(s)
	return nil
}
