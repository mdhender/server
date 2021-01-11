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
	"strings"
)

type State struct {
	turn     int
	admins   map[string]bool // id of the administrator
	polities map[string]*Polity
	systems  map[string]*System
	stars    map[string]*Star
	planets  map[string]*Planet
	colonies map[string]*Colony
	ships    map[string]*Ship

	orders Orders
}

// NewState returns an initialized state with an administrator.
func NewState(admins ...string) (*State, error) {
	cluster := mkcluster()
	st := &State{
		admins:   cluster.admins,
		polities: cluster.polities,
		systems:  cluster.systems,
		stars:    cluster.stars,
		planets:  cluster.planets,
		colonies: cluster.colonies,
		ships:    cluster.ships,
	}

	if len(admins) == 0 {
		// add the default administrator id
		st.admins[uuid.New().String()] = true
	} else {
		// add names that the caller passed in
		for _, admin := range admins {
			if admin != strings.TrimSpace(sanitize(admin)) {
				return nil, fmt.Errorf("invalid characters in admin: %w", ERRBADREQUEST)
			}
			st.admins[admin] = true
		}
	}

	// and return it all
	return st, nil
}

func (st *State) Admins() []string {
	var s []string
	for id := range st.admins {
		s = append(s, id)
	}
	return s
}

func (st *State) Colony(id string) *Colony {
	if c, ok := st.colonies[id]; ok {
		return c
	}
	return nil
}

func (st *State) ProcessOrders(orders Orders, debug bool) []error {
	st.turn++
	return st.ExecuteOrders(orders, debug)
}

func (st *State) Planet(id string) *Planet {
	if p, ok := st.planets[id]; ok {
		return p
	}
	return nil
}

func (st *State) Polity(id string) *Polity {
	if p, ok := st.polities[id]; ok {
		return p
	}
	return nil
}

func (st *State) Ship(id string) *Ship {
	if s, ok := st.ships[id]; ok {
		return s
	}
	return nil
}

func (st *State) Star(id string) *Star {
	if s, ok := st.stars[id]; ok {
		return s
	}
	return nil
}

func (st *State) System(id string) *System {
	if s, ok := st.systems[id]; ok {
		return s
	}
	return nil
}

type AutomationUnit struct{}
type EngineUnit struct{}
type FactoryUnit struct{}
type FarmUnit struct {
	techLevel int
	quantity  int
}

func (u FarmUnit) Produce() int {
	if u.techLevel == 1 {
		return 25 * u.quantity
	}
	return 5 * u.techLevel * u.quantity
}

type MineUnit struct {
	techLevel int
	quantity  int
	resource  *Resource // resource being mined
}
type MissileUnit struct{}
type AntiMissileUnit struct{}

type RobotUnit struct{}

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
	return c.polity == p
}

// acceptsOrdersFrom implements the hegemony interface.
func (s *Ship) acceptsOrdersFrom(p *Polity) bool {
	if s == nil {
		return false
	}
	return s.polity == p
}

func (st *State) assignColonyName(colony *Colony, name string) error {
	if colony == nil {
		return fmt.Errorf("missing colony: %w", ERRBADREQUEST)
	}
	if name = strings.TrimSpace(name); name == "" || len(name) > 50 {
		return fmt.Errorf("invalid name %q: %w", name, ERRBADREQUEST)
	}
	return fmt.Errorf("State.assignColonyName: %w", ERRNOTIMPLEMENTED)
}

func (st *State) assignColonyNote(colony *Colony, note Text) error {
	if colony == nil {
		return fmt.Errorf("missing colony: %w", ERRBADREQUEST)
	}
	if note = note.TrimSpace(); len(note.text) > 200 {
		return fmt.Errorf("invalid note: %w", ERRBADREQUEST)
	}
	colony.note = note
	return nil
}

func (st *State) assignPlanetName(planet *Planet, name string) error {
	if planet == nil {
		return fmt.Errorf("missing planet: %w", ERRBADREQUEST)
	}
	if name = strings.TrimSpace(name); name == "" || len(name) > 50 {
		return fmt.Errorf("invalid name %q: %w", name, ERRBADREQUEST)
	}
	return fmt.Errorf("State.assignPlanetName: %w", ERRNOTIMPLEMENTED)
}

func (st *State) assignPolityName(polity *Polity, name string) error {
	if polity == nil {
		return fmt.Errorf("missing polity: %w", ERRBADREQUEST)
	}
	if name = strings.TrimSpace(name); name == "" || len(name) > 50 {
		return fmt.Errorf("invalid name %q: %w", name, ERRBADREQUEST)
	}
	return fmt.Errorf("State.assignPolityName: %w", ERRNOTIMPLEMENTED)
}

func (st *State) assignShipName(ship *Ship, name string) error {
	if ship == nil {
		return fmt.Errorf("missing ship: %w", ERRBADREQUEST)
	}
	if name = strings.TrimSpace(name); name == "" || len(name) > 50 {
		return fmt.Errorf("invalid name %q: %w", name, ERRBADREQUEST)
	}
	return fmt.Errorf("State.assignShipName: %w", ERRNOTIMPLEMENTED)
}

func (st *State) assignShipNote(ship *Ship, note Text) error {
	if ship == nil {
		return fmt.Errorf("missing ship: %w", ERRBADREQUEST)
	}
	if note = note.TrimSpace(); len(note.text) > 200 {
		return fmt.Errorf("invalid note: %w", ERRBADREQUEST)
	}
	ship.note = note
	return nil
}

func (st *State) assignStarName(star *Star, name string) error {
	if star == nil {
		return fmt.Errorf("missing star: %w", ERRBADREQUEST)
	}
	if name = strings.TrimSpace(name); name == "" || len(name) > 50 {
		return fmt.Errorf("invalid name %q: %w", name, ERRBADREQUEST)
	}
	return fmt.Errorf("State.assignStarName: %w", ERRNOTIMPLEMENTED)
}

func (st *State) assignSystemName(system *System, name string) error {
	if system == nil {
		return fmt.Errorf("missing system: %w", ERRBADREQUEST)
	}
	if name = strings.TrimSpace(name); name == "" || len(name) > 50 {
		return fmt.Errorf("invalid name %q: %w", name, ERRBADREQUEST)
	}
	return fmt.Errorf("State.assignSystemName: %w", ERRNOTIMPLEMENTED)
}

// isDuplicateID returns true if the id is already in a map.
func (st *State) isDuplicateID(id string) bool {
	if _, ok := st.colonies[id]; ok {
		return true
	} else if _, ok := st.planets[id]; ok {
		return true
	} else if _, ok := st.polities[id]; ok {
		return true
	} else if _, ok := st.ships[id]; ok {
		return true
	} else if _, ok := st.stars[id]; ok {
		return true
	} else if _, ok := st.systems[id]; ok {
		return true
	}
	return false
}

// permissionToColonize enables a ship to establish a new colony on a planet.
// Permission expires at the end of the current turn.
func (st *State) permissionToColonize(planet *Planet, ship *Ship) error {
	if planet == nil {
		return fmt.Errorf("missing planet: %w", ERRBADREQUEST)
	} else if ship == nil {
		return fmt.Errorf("missing ship: %w", ERRBADREQUEST)
	}

	return fmt.Errorf("State.permissionToColonize: %w", ERRNOTIMPLEMENTED)
}

// setHomePort updates the home port of the given ship
func (st *State) setHomePort(ship *Ship, colony *Colony) error {
	if ship == nil {
		return fmt.Errorf("missing ship: %w", ERRBADREQUEST)
	} else if colony == nil {
		return fmt.Errorf("missing colony: %w", ERRBADREQUEST)
	} else if ship.polity != colony.polity {
		return fmt.Errorf("ship not aligned to colony: %w", ERRBADREQUEST)
	} else if ship.homePort == colony {
		return nil // already assigned, so nothing to do
	}
	return colony.xferShip(ship)
}

// transferColony transfers control of a colony to another Polity.
// Can be used to give away or seize control of a colony.
func (st *State) transferColony(colony *Colony, from, to *Polity) error {
	if to == nil {
		return fmt.Errorf("missing to: %w", ERRBADREQUEST)
	} else if colony == nil {
		return fmt.Errorf("missing colony: %w", ERRBADREQUEST)
	} else if colony.polity == to {
		return nil // nothing to do
	}
	return to.xferColony(colony)
}

// transferPolity transfers control of a Polity to another Polity.
// Can be used when a new player joins a game or when a player exits the game.
//
// When transferring control, the diplomatic status is purged, making
// the new controller the ally of the Polity.
func (st *State) transferPolity(from, to *Polity) error {
	if from == nil {
		return fmt.Errorf("missing from: %w", ERRBADREQUEST)
	} else if to == nil {
		return fmt.Errorf("missing to: %w", ERRBADREQUEST)
	} else if from == to {
		return nil // nothing to do
	}

	from.diplomacy = make(map[string]DiplomaticStatus)
	for _, polity := range st.polities {
		from.diplomacy[polity.id] = UNKNOWN
	}
	from.diplomacy[to.id] = ALLY
	if to.diplomacy == nil {
		to.diplomacy = make(map[string]DiplomaticStatus)
	}
	to.diplomacy[from.id] = ALLY

	return nil
}

// transferShip transfers control of a ship to another Polity.
// Can be used to give away or seize control of a ship.
func (st *State) transferShip(ship *Ship, from, to *Polity) error {
	if to == nil {
		return fmt.Errorf("missing to: %w", ERRBADREQUEST)
	} else if ship == nil {
		return fmt.Errorf("missing ship: %w", ERRBADREQUEST)
	} else if ship.polity == to {
		return nil // nothing to do
	}
	return to.xferShip(ship)
}
