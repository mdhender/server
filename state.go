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

import "log"

type State struct {
	maps struct {
		// ids is a global hot mess
		ids map[string]interface{}
	}
	polities []*Polity
}

func NewState() *State {
	st := &State{}
	st.maps.ids = make(map[string]interface{})
	return st
}

func (st *State) Lookup(id string) (o interface{}, ok bool) {
	if st.maps.ids != nil {
		o, ok = st.maps.ids[id]
	}
	return o, ok
}

func (st *State) LookupColony(id string) *Colony {
	if o, ok := st.Lookup(id); ok {
		if colony, ok := o.(*Colony); ok {
			return colony
		}
	}
	return nil
}

func (st *State) LookupPolity(id string) *Polity {
	if o, ok := st.Lookup(id); ok {
		if p, ok := o.(*Polity); ok {
			return p
		}
	}
	return nil
}

func (st *State) LookupShip(id string) *Ship {
	if o, ok := st.Lookup(id); ok {
		if ship, ok := o.(*Ship); ok {
			return ship
		}
	}
	return nil
}

type System struct {
	ID     string
	Name   string
	Coords struct {
		X int
		Y int
		Z int
	}
	Stars []*Star
}

// I think that orbit 11 is the default jump point target.
type Star struct {
	ID     string
	Name   string
	Orbits [11]*Orbit
}

type Orbit struct {
	ID       string
	Planet   *Planet
	Colonies []*Colony
	Ships    []*Ship
}

type Planet struct {
	ID           string
	Type         PlanetType
	Habitability int // range from 0 to 25
	Deposits     []*Resource
}

type Polity struct {
	ID   string
	Name string
	Home struct {
		System *System
		World  *Planet
	}
	// controls and controlled-by allow the players to join and
	// quit games in progress. there should never be more than
	// one level in this hierarchy.
	controls  []*Polity // controls these polities
	viceroyOf *Polity   // controlled by this polity
	diplomacy map[string]DiplomaticStatus
}

func (p *Polity) isAllied(t *Polity) bool {
	if p == nil || t == nil {
		return false
	}
	return p.diplomacy[t.ID] == ALLY && t.diplomacy[p.ID] == ALLY
}

// Resource is any resource that can be mined.
type Resource struct {
	ID              string
	Type            ResourceType
	YieldPct        float64
	Unlimited       bool
	InitialAmount   int64
	AmountRemaining int64
}

type AutomationUnit struct{}
type Colony struct {
	originalPolity *Polity // set only if this is a Home Colony?
	controlledBy   *Polity
	system         *System
	homePortTo     []*Ship
}

type EngineUnit struct{}
type FactoryUnit struct{}
type FarmUnit struct{}
type MineUnit struct{}
type MissileUnit struct{}
type AntiMissileUnit struct{}
type PopulationUnit struct{}
type RobotUnit struct{}

type Ship struct {
	controlledBy *Polity
	system       *System
	homePort     *Colony
}

type StructuralUnit struct{}
type TransportUnit struct{}

// hegemony defines the contract for accepting orders from either a ruler or a viceroy of the ruler.
type hegemony interface {
	// acceptsOrdersFrom returns true if the unit is controlled directly or indirectly by the ruler.
	acceptsOrdersFrom(p *Polity) bool
}

// acceptsOrdersFrom implements the hegemony interface.
func (c *Colony) acceptsOrdersFrom(p *Polity) bool {
	if c == nil {
		return false
	}
	for _, viceroy := range p.controls {
		if c.controlledBy == viceroy {
			return true
		}
	}
	return c.controlledBy == p
}

// acceptsOrdersFrom implements the hegemony interface.
func (s *Ship) acceptsOrdersFrom(p *Polity) bool {
	if s == nil {
		return false
	}
	for _, viceroy := range p.controls {
		if s.controlledBy == viceroy {
			return true
		}
	}
	return s.controlledBy == p
}

// assignHomePort updates the home port of the given ship
func (st *State) assignHomePort(ship *Ship, colony *Colony) error {
	if ship == nil {
		log.Printf("[bug] updateHomePort: ship is nil\n")
		return ERRBUG
	} else if colony == nil {
		log.Printf("[bug] updateHomePort: colony is nil\n")
		return ERRBUG
	} else if ship.controlledBy != colony.controlledBy {
		log.Printf("[bug] updateHomePort: assert(ship aligned to colony)\n")
		return ERRBUG
	} else if ship.homePort == colony {
		// already assigned, so nothing to do
		return nil
	}

	// remove the ship from its current home port
	if ship.homePort != nil {
		var ships []*Ship
		for _, s := range ship.homePort.homePortTo {
			if s != ship {
				ships = append(ships, s)
			}
		}
		if len(ships) == len(ship.homePort.homePortTo) {
			log.Printf("[bug] updateHomePort: assert(ship's prior homePort valid)\n")
		}
		ship.homePort.homePortTo = ships
	}

	ship.homePort = colony
	colony.homePortTo = append(colony.homePortTo, ship)

	return nil
}

// transferColony transfers control of a colony to another Polity.
// Can be used to give away or seize control of a colony.
func (st *State) transferColony(colony *Colony, from, to *Polity) error {
	if colony == nil {
		log.Printf("[bug] transferColony: colony is nil\n")
		return ERRBUG
	} else if to == nil {
		log.Printf("[bug] transferColony: to is nil\n")
		return ERRBUG
	}
	// remove from current controller
	log.Printf("[todo] transferColony: remove from current controller\n")
	// add to new controller.
	colony.controlledBy = to
	// recursively transfer control of all assets assigned to the colony.
	return nil
}

// transferPolity transfers control of a Polity to another Polity.
// Can be used when a new player joins a game or when a player exits the game.
//
// When transferring control, the diplomatic status is purged, making
// the new controller the ally of the Polity.
func (st *State) transferPolity(from, to *Polity) error {
	if from == nil {
		log.Printf("[bug] transferPolity: from is nil\n")
		return ERRBUG
	} else if to == nil {
		log.Printf("[bug] transferPolity: to is nil\n")
		return ERRBUG
	} else if from == to {
		log.Printf("[bug] transferPolity: from == to\n")
		return ERRBUG
	}

	from.diplomacy = make(map[string]DiplomaticStatus)
	for _, polity := range st.polities {
		from.diplomacy[polity.ID] = UNKNOWN
	}
	from.diplomacy[to.ID] = ALLY
	if to.diplomacy == nil {
		to.diplomacy = make(map[string]DiplomaticStatus)
	}
	to.diplomacy[from.ID] = ALLY

	return nil
}

// transferShip transfers control of a ship to another Polity.
// Can be used to give away or seize control of a ship.
func (st *State) transferShip(ship *Ship, from, to *Polity) error {
	if ship == nil {
		log.Printf("[bug] transferShip: ship is nil\n")
		return ERRBUG
	} else if to == nil {
		log.Printf("[bug] transferShip: to is nil\n")
		return ERRBUG
	}
	// remove from current controller
	log.Printf("[todo] transferShip: remove from current controller\n")
	// add to new controller
	ship.controlledBy = to
	// recursively transfer control of all assets assigned to the ship.
	return nil
}

// diplomaticStatus returns the status that the polity thinks that it
// has with the target.
func (p *Polity) diplomaticStatus(target *Polity) DiplomaticStatus {
	if p.diplomacy != nil {
		if ds, ok := p.diplomacy[target.ID]; ok {
			return ds
		}
	}
	return UNKNOWN
}

// controlledBy returns the entity that controls the polity.
// In less confusing terms, if the polity is a viceroy, we return
// the polity we are controlled by. otherwise, we return ourself.
func (p *Polity) controlledBy() *Polity {
	if p == nil || p.viceroyOf == nil {
		return p
	}
	return p.viceroyOf
}

func (c *Colony) isHomeColony() bool {
	panic(ERRNOTIMPLEMENTED)
}

func (p *Polity) isViceroyOf(ruler *Polity) bool {
	if p == nil || ruler == nil {
		return false
	}
	if p.viceroyOf == ruler {
		for _, pp := range ruler.controls {
			if pp == p {
				return true
			}
		}
		log.Printf("[bug] Polity.isViceroy: assert(ruler.controls contains p.viceroyOf)\n")
		return false
	}
	for _, pp := range ruler.controls {
		if pp == p {
			log.Printf("[bug] Polity.isViceroy: assert(ruler.controls does not contain p.viceroyOf)\n")
			return false
		}
	}
	return false
}
