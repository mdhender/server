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

import (
	"fmt"
	"github.com/google/uuid"
	"log"
)

// notes on units...
// 18_446_744_073_709_551_615 // maximum unsigned 64 bit integer
//              4_294_967_295 // maximum unsigned 32 bit integer
//                 10_000_000 // people in one population unit
//                 40_000_000 // people fed by one food unit per turn
//              1_000_000_000 // people fed by FARM-1 per turn

func NewState() (*State, error) {
	st := &State{}
	st.maps.ids = make(map[string]interface{})

	for _, input := range []*struct {
		name    string
		sysname string
		x, y, z int
	}{
		{"usagi", "Shikoku", 1, 1, 1},
		{"tomoe", "Kyushu", 2, 2, 2},
	} {
		polity := st.MakePolity(input.name)
		log.Printf("[state] polity %q\n", polity.name)

		system := st.MakeSystem(input.sysname, input.x, input.y, input.z)
		log.Printf("[state] system %q\n", system.name)

		system.stars[0].orbits[5].planet = st.MakeHomePlanet(polity)
	}

	return st, nil
}

func (st *State) MakeColony(kind ColonyKind, polity *Polity) *Colony {
	colony := &Colony{id: uuid.New().String(), kind: kind, controlledBy: polity}
	st.maps.ids[colony.id] = colony
	colony.number = colony.controlledBy.nextColonyNumber()
	polity.controls.colonies = append(polity.controls.colonies)
	return colony
}

func (st *State) MakeHomePlanet(polity *Polity, scarcity bool) *Planet {
	planet := &Planet{id: uuid.New().String()}
	st.maps.ids[planet.id] = planet
	planet.kind = TERRESTRIAL
	planet.habitability = 25 // in tens of millions

	homeColony := st.MakeColony(OPEN, polity)
	homeColony.originalPolity = polity
	homeColony.population.construction = 20_000
	homeColony.population.professionals = 2_000_000
	homeColony.population.soldiers = 2_500_000
	homeColony.population.unskilled = 6_000_000
	homeColony.population.others = 5_900_000
	homeColony.population.total = homeColony.population.construction + homeColony.population.professionals + homeColony.population.soldiers + homeColony.population.spies + homeColony.population.trainees + homeColony.population.unskilled + homeColony.population.others
	homeColony.storage.fuel = 2_000_000
	homeColony.storage.gold = 50_000
	homeColony.storage.metal = 4_000_000
	homeColony.storage.nonmetal = 4_000_000
	min, max := homeColony.population.FoodNeeded()
	if scarcity {// put the minimum amount of food needed into storage
		homeColony.storage.food = min
	} else {
		homeColony.storage.food = max
	}

	for _, kind := range []ResourceKind{FUEL, GOLD, METAL, NONMETAL} {
		mine := MineUnit{techLevel: 1, quantity: 1, resource: st.MakeResource(kind, true)}
		planet.deposits = append(planet.deposits, mine.resource)
		homeColony.units.mines = append(homeColony.units.mines, mine)
	}

	// every home world gets an orbiting colony
	orbitingColony := st.MakeColony(ORBITING, polity)
	orbitingColony.population.construction = 10_000
	orbitingColony.population.professionals = 100_000
	orbitingColony.population.soldiers = 150_000
	orbitingColony.population.unskilled = 370_000
	orbitingColony.population.others = 350_000
	orbitingColony.population.total = orbitingColony.population.construction + orbitingColony.population.professionals + orbitingColony.population.soldiers + orbitingColony.population.spies + orbitingColony.population.trainees + orbitingColony.population.unskilled + orbitingColony.population.others
	orbitingColony.storage.fuel = 200_000
	orbitingColony.storage.gold = 5_000
	orbitingColony.storage.metal = 400_000
	orbitingColony.storage.nonmetal = 400_000
	min, max = orbitingColony.population.FoodNeeded()
	if scarcity {// put the minimum amount of food needed into storage
		orbitingColony.storage.food = min
	} else {
		orbitingColony.storage.food = max
	}

	// need enough farms to produce enough food to sustain the population on the planet and in orbit.
	// scarcity impacts the total number of farms because it is based on the amount allocated to storage.
	// FARM-1 produces 25 Food Units per turn; provide a bit of room for growth
	farmUnitsNeeded := ((homeColony.storage.food+orbitingColony.storage.food) / 25) + 1
	homeColony.units.farms = append(homeColony.units.farms, FarmUnit{techLevel: 1, quantity: farmUnitsNeeded})

	polity.home.world = planet
	polity.home.colony = homeColony
	return planet
}

func (st *State) MakePolity(name string) *Polity {
	p := &Polity{id: uuid.New().String()}
	st.maps.ids[p.id] = p
	p.name = name
	return p
}

func (st *State) MakeResource(kind ResourceKind, unlimited bool) *Resource {
	resource := &Resource{id: uuid.New().String()}
	st.maps.ids[resource.id] = resource
	resource.kind = kind
	resource.unlimited = unlimited
	if resource.unlimited {
		resource.yieldPct = 0.03
	} else {
		if resource.kind == GOLD {
			resource.yieldPct = 0.09
		} else {
			resource.yieldPct = 0.90
		}
		resource.initialAmount = 55 * 1_000_000_000
		resource.amountRemaining = resource.initialAmount
	}
	return resource
}

func (st *State) MakeSystem(name string, x, y, z int) *System {
	s := &System{id: uuid.New().String()}
	st.maps.ids[s.id] = s
	s.coords.x, s.coords.y, s.coords.z = x, y, z
	s.name = fmt.Sprintf("%02d-%02d-%02d", x, y, z)
	star := st.MakeStar(s.name)
	s.stars = append(s.stars, star)
	return s
}

func (st State) MakeStar(name string) *Star {
	star := &Star{id: uuid.New().String()}
	st.maps.ids[star.id] = star
	star.name = name
	for orbit := 0; orbit < 11; orbit++ {
		star.orbits[orbit] = &Orbit{}
	}
	return star
}

func (star *Star) AddHomeWorld(planet *Planet) {
	star.orbits[5].planet = planet
}
