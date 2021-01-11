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
	"log"
)

type Polity struct {
	id   string
	name string
	home struct {
		system *System
		star   *Star
		planet *Planet
		colony *Colony
		world  string
	}
	controls struct {
		colonies map[string]*Colony
		polities map[string]*Polity // there should never be more than one level in this hierarchy.
		ships    map[string]*Ship
	}
	viceroyOf *Polity // viceroy to this polity
	diplomacy map[string]DiplomaticStatus
	seq       struct {
		colony int
		ship   int
	}
}

func polity() *Polity {
	p := &Polity{id: uuid.New().String()}
	p.controls.colonies = make(map[string]*Colony)
	p.controls.polities = make(map[string]*Polity)
	p.controls.ships = make(map[string]*Ship)
	p.diplomacy = make(map[string]DiplomaticStatus)
	return p
}

func (p *Polity) addColony(c *Colony) {
	if p == nil || c == nil {
		return
	}
	p.controls.colonies[c.id] = c
}

func (p *Polity) addShip(s *Ship) {
	if p == nil || s == nil {
		return
	}
	p.controls.ships[s.id] = s
}

func (p *Polity) delColony(c *Colony) {
	if p == nil || c == nil {
		return
	}
	delete(p.controls.colonies, c.id)
}

func (p *Polity) delShip(s *Ship) {
	if p == nil || s == nil {
		return
	}
	delete(p.controls.ships, s.id)
}

// diplomaticStatus returns the status that the polity thinks that it
// has with the target.
func (p *Polity) diplomaticStatus(t *Polity) DiplomaticStatus {
	if p == nil || p.diplomacy == nil || t == nil {
		return UNKNOWN
	}
	if ds, ok := p.diplomacy[t.id]; ok {
		return ds
	}
	return UNKNOWN
}

func (p *Polity) isAlliedTo(t *Polity) bool {
	return p.isViceroyOf(t) || (p.diplomaticStatus(t) == ALLY && t.diplomaticStatus(p) == ALLY)
}

func (p *Polity) isAllyOf(t *Polity) bool {
	return p.diplomaticStatus(t) == ALLY
}

func (p *Polity) isViceroyOf(t *Polity) bool {
	if p == nil || p.viceroyOf == nil {
		return false
	}
	return p.viceroyOf == t
}

func (p *Polity) nextColonyNumber() string {
	number := fmt.Sprintf("C%d", p.seq.colony)
	p.seq.colony++
	return number
}

func (p *Polity) nextShipNumber() string {
	p.seq.ship++
	return fmt.Sprintf("C%d", p.seq.ship)
}

func (p *Polity) xferColony(c *Colony) error {
	if p == nil || c == nil {
		return nil
	}
	c.polity.delColony(c)
	p.addColony(c)
	// recursively transfer control of all assets assigned to the colony.
	log.Printf("[todo] xferColony: recursively transfer control of all assets assigned to the colony\n")
	return nil
}

func (p *Polity) xferShip(s *Ship) error {
	if p == nil || s == nil {
		return nil
	}
	s.polity.delShip(s)
	p.addShip(s)
	// recursively transfer control of all assets assigned to the ship.
	log.Printf("[todo] xferShip: recursively transfer control of all assets assigned to the ship\n")
	return nil
}
