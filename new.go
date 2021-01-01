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
	colony := &Colony{id: uuid.New().String(), controlledBy: polity}
	st.maps.ids[colony.id] = colony
	colony.number = colony.controlledBy.nextColonyNumber()
	polity.controls.colonies = append(polity.controls.colonies)
	return colony
}

func (st *State) MakeHomePlanet(polity *Polity) *Planet {
	planet := &Planet{id: uuid.New().String()}
	st.maps.ids[planet.id] = planet
	planet.kind = TERRESTRIAL
	planet.habitability = 25 // in tens of millions

	// 18_446_744_073_709_551_615 // maximum unsigned 64 bit integer
	//              4_294_967_295 // maximum unsigned 32 bit integer
	//                 10_000_000 // people in one population unit

	colony := st.MakeColony(OPEN, polity)
	colony.originalPolity = polity
	colony.units.population.construction = 20_000
	colony.units.population.professionals = 2_000_000
	colony.units.population.soldiers = 2_500_000
	colony.units.population.unskilled = 6_000_000
	colony.units.population.others = 5_900_000
	colony.storage.food = 3_000_000
	colony.storage.fuel = 2_000_000
	colony.storage.gold = 50_000
	colony.storage.metal = 4_000_000
	colony.storage.nonmetal = 4_000_000

	orbitingColony := st.MakeColony(ORBITING, polity)
	orbitingColony.units.population.construction = 10_000
	orbitingColony.units.population.professionals = 100_000
	orbitingColony.units.population.soldiers = 150_000
	orbitingColony.units.population.unskilled = 370_000
	orbitingColony.units.population.others = 350_000
	orbitingColony.storage.food = 300_000
	orbitingColony.storage.fuel = 200_000
	orbitingColony.storage.gold = 5_000
	orbitingColony.storage.metal = 400_000
	orbitingColony.storage.nonmetal = 400_000

	var initialPopulation int
	initialPopulation += colony.units.population.construction
	initialPopulation += colony.units.population.professionals
	initialPopulation += colony.units.population.soldiers
	initialPopulation += colony.units.population.spies
	initialPopulation += colony.units.population.trainees
	initialPopulation += colony.units.population.unskilled
	initialPopulation += colony.units.population.others
	initialPopulation += orbitingColony.units.population.construction
	initialPopulation += orbitingColony.units.population.professionals
	initialPopulation += orbitingColony.units.population.soldiers
	initialPopulation += orbitingColony.units.population.spies
	initialPopulation += orbitingColony.units.population.trainees
	initialPopulation += orbitingColony.units.population.unskilled
	initialPopulation += orbitingColony.units.population.others

	// need enough farms to sustain the population.
	// one population unit is 10_000_000 people.
	populationUnits := (initialPopulation / 10_000_000) + 1
	// Each Population Unit consumes 0.25 Food Units per turn.
	foodUnitsNeeded := (populationUnits / 4) + 1
	// FARM-1 produces 25 Food Units per turn, which is enough to feed 100 population units
	farmUnitsNeeded := (foodUnitsNeeded / 25) + 1
	colony.units.farms = append(colony.units.farms, FarmUnit{techLevel: 1, quantity: farmUnitsNeeded})

	for _, kind := range []ResourceKind{FUEL, GOLD, METAL, NONMETAL} {
		mine := MineUnit{techLevel: 1, quantity: 1, resource: st.MakeResource(kind, true)}
		planet.deposits = append(planet.deposits, mine.resource)
		colony.units.mines = append(colony.units.mines, mine)
	}

	polity.home.world = planet
	polity.home.colony = colony
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
