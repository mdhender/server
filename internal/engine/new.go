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

//func NewState() (*State, error) {
//	st := &State{}
//	st.ids = make(map[string]interface{})
//
//	var scarcity bool
//	for _, input := range []*struct {
//		name    string
//		sysname string
//		x, y, z int
//	}{
//		{"usagi", "Shikoku", 1, 1, 1},
//		{"tomoe", "Kyushu", 2, 2, 2},
//	} {
//		polity := st.MakePolity(input.name)
//		log.Printf("[state] polity %q\n", polity.name)
//
//		system := st.MakeSystem(input.sysname, input.x, input.y, input.z)
//		log.Printf("[state] system %q\n", system.name)
//		polity.home.system = system
//
//		system.stars[0].orbits[5].planet = st.MakeHomePlanet(polity, scarcity)
//
//		scarcity = !scarcity
//	}
//
//	return st, nil
//}
//
//func (st *State) MakeHomePlanet(polity *Polity, scarcity bool) string {
//	planet := Planet{id: uuid.New().String()}
//	planet.kind = TERRESTRIAL
//	planet.habitability = 25 // in tens of millions
//
//	pop := population.NewHomeColony(false)
//	homeColony := colony.New(colony.OPEN, polity.id, polity.nextColonyNumber(), pop)
//	polity.controls.colonies[homeColony.ID()] = homeColony
//	planet.colonies = append(planet.colonies, homeColony.ID())
//
//	if scarcity {
//		homeColony.Ration = 0.25
//	}
//	upop := units.Unit{Kind: units.POPULATION, TechLevel: 1, Quantity: pop.TotalCount()}
//	homeColony.Units = append(homeColony.Units, upop)
//
//	// determine storage needed for the initial population
//	volume := upop.Volume()
//
//	// add initial resources and associated storage
//	for _, r := range []struct {
//		kind units.Kind
//		qty  int
//	}{
//		{units.FUEL, 2_000_000},
//		{units.GOLD, 50_000},
//		{units.METAL, 4_000_000},
//		{units.NONMETAL, 4_000_000},
//	} {
//		u := units.Unit{Kind: r.kind, Quantity: r.qty}
//		homeColony.Units = append(homeColony.Units, u)
//		volume += u.Volume()
//		fmt.Printf("[todo] homeColony.stockpile            = %s\n", u.Sexpr())
//	}
//
//	// the population on all colonies want to stockpile a certain amount of food
//	foodMin, foodMax := pop.FoodNeededPerTurn()
//	homeColony.Stockpiles.Food.Want = foodMax * 4
//	if scarcity { // put the minimum amount of food needed into storage
//		homeColony.Stockpiles.Food.Have = foodMin
//	} else {
//		homeColony.Stockpiles.Food.Have = foodMax
//	}
//	// create storage for the stockpiles
//	foodUnits := units.Unit{Kind: units.FOOD, TechLevel: 1, Quantity: homeColony.Stockpiles.Food.Have}
//	homeColony.Units = append(homeColony.Units, foodUnits)
//
//	// update volume of storage needed
//	volume += foodUnits.Volume()
//
//	// determine number of farm units needed to produce that amount of food each turn
//	fmt.Printf("[todo] homeColony.population           = %13s\n", utils.Commas(pop.TotalCount()))
//	fmt.Printf("[todo] homeColony.food.neededPerTurn   = %13s\n", utils.Commas(foodMax))
//	farmUnits := units.Unit{Kind: units.FARM, Assembled: true, TechLevel: 1, Quantity: 1}
//	fmt.Printf("[todo] homeColony.food.producedPerTurn = %13s\n", utils.Commas(farmUnits.Produce().Quantity))
//	farmUnits.Quantity = (foodMax / farmUnits.Produce().Quantity) + 1
//	fmt.Printf("[todo] homeColony.farmUnits.Quantity   = %13s\n", utils.Commas(farmUnits.Quantity))
//	fmt.Printf("[todo] homeColony.farmUnits.Produce    = %13s\n", utils.Commas(farmUnits.Produce().Quantity))
//	homeColony.Units = append(homeColony.Units, farmUnits)
//
//	for _, kind := range []ResourceKind{FUEL, GOLD, METAL, NONMETAL} {
//		resource := st.MakeResource(kind, true)
//		planet.deposits = append(planet.deposits, resource)
//		mine := units.Unit{Kind: units.MINE, TechLevel: 1, Quantity: 100_000}
//		homeColony.Units = append(homeColony.Units, mine)
//	}
//
//	// add hydro-electric power plants
//	powerUnits := units.Unit{Kind: units.POWER, Assembled: true, TechLevel: 1, Quantity: 100_000}
//	homeColony.Units = append(homeColony.Units, powerUnits)
//
//	// bump the volume by by 25% to give the colony some room to expand,
//	// then add structural units needed to enclose that volume to the colony.
//	structuralUnitsNeeded := int(float64(volume)*1.25) * homeColony.StructureFactor()
//	homeColony.Units = append(homeColony.Units, units.Unit{Kind: units.STRUCTURAL, Assembled: true, TechLevel: 1, Quantity: structuralUnitsNeeded})
//
//	st.ids[planet.id] = planet
//	st.ids[homeColony.ID()] = homeColony
//
//	//// every home world gets an orbiting colony
//	//pop = population.NewHomeColony(true)
//	//orbitingColony := colony.New(colony.OPEN, polity.nextColonyNumber(), pop)
//	//st.ids[orbitingColony.ID()] = orbitingColony
//	//polity.controls.colonies = append(polity.controls.colonies, orbitingColony.ID())
//	//if scarcity {
//	//	orbitingColony.Ration = 0.25
//	//}
//	//
//	//fmt.Printf("[todo] orbitingColony.storage.fuel = 200_000")
//	//fmt.Printf("[todo] orbitingColony.storage.gold = 5_000")
//	//fmt.Printf("[todo] orbitingColony.storage.metal = 400_000")
//	//fmt.Printf("[todo] orbitingColony.storage.nonmetal = 400_000")
//	//
//	//// the population on all colonies want to stockpile a certain amount of food and goods
//	//min, max = pop.FoodNeededPerTurn()
//	//orbitingColony.Stockpiles.Food.Want = max
//	//if scarcity { // put the minimum amount of food needed into storage
//	//	orbitingColony.Stockpiles.Food.Have = max
//	//} else {
//	//	orbitingColony.Stockpiles.Food.Have = min
//	//}
//	//// create storage for the stockpiles
//	//foodUnits = units.Unit{Kind: units.FOOD, Quantity: homeColony.Stockpiles.Food.Want}
//	//storageUnits = units.Unit{Kind: units.STRUCTURAL, Quantity: foodUnits.Volume()}
//	//
//	//// need enough farms to produce enough food to sustain the population on the planet and in orbit.
//	//// scarcity impacts the total number of farms because it is based on the amount allocated to storage.
//	//// FARM-1 produces 25 Food Units per turn; provide a bit of room for growth
//	//farmUnitsNeeded := ((orbitingColony.storage.food + orbitingColony.storage.food) / 25) + 1
//	//orbitingColony.units.farms = append(orbitingColony.units.farms, FarmUnit{techLevel: 1, quantity: farmUnitsNeeded})
//	//
//	//// bump the number of storage units to give the colony some room to expand
//	//storageUnits.Quantity = int(float64(storageUnits.Quantity) * 1.25)
//	//// then add them to the colony.
//	//orbitingColony.Units = append(orbitingColony.Units, storageUnits)
//
//	polity.home.world = planet.id
//	polity.home.colony = homeColony.ID()
//
//	return planet.id
//}
//
//func (st *State) MakePolity(name string) *Polity {
//	p := &Polity{id: uuid.New().String()}
//	st.ids[p.id] = p
//	st.polities = append(st.polities, p)
//	p.name = name
//	return p
//}
//
//func (st *State) MakeResource(kind ResourceKind, unlimited bool) *Resource {
//	resource := &Resource{id: uuid.New().String()}
//	st.ids[resource.id] = resource
//	resource.kind = kind
//	resource.unlimited = unlimited
//	if resource.unlimited {
//		resource.yieldPct = 0.03
//	} else {
//		if resource.kind == GOLD {
//			resource.yieldPct = 0.09
//		} else {
//			resource.yieldPct = 0.90
//		}
//		resource.initialAmount = 55 * 1_000_000_000
//		resource.amountRemaining = resource.initialAmount
//	}
//	return resource
//}
//
//func (st *State) MakeSystem(name string, x, y, z int) *System {
//	s := &System{id: uuid.New().String()}
//	st.ids[s.id] = s
//	s.coords.x, s.coords.y, s.coords.z = x, y, z
//	s.name = fmt.Sprintf("%02d-%02d-%02d", x, y, z)
//	star := st.MakeStar(s.name)
//	s.stars = append(s.stars, star)
//	return s
//}
//
//func (st State) MakeStar(name string) *Star {
//	star := &Star{id: uuid.New().String()}
//	st.ids[star.id] = star
//	star.name = name
//	for orbit := 0; orbit < len(star.orbits); orbit++ {
//		star.orbits[orbit] = &Orbit{}
//	}
//	return star
//}
//
//func (star *Star) AddHomeWorld(id string) {
//	star.orbits[5].planet = id
//}
