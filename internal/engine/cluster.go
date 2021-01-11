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
)

type Cluster struct {
	admins   map[string]bool // id of the administrator
	colonies map[string]*Colony
	planets  map[string]*Planet
	polities map[string]*Polity
	ships    map[string]*Ship
	stars    map[string]*Star
	systems  map[string]*System // system and star system are the same
}

func mkcluster() *Cluster {
	cluster := &Cluster{
		admins:   make(map[string]bool),
		polities: make(map[string]*Polity),
		systems:  make(map[string]*System),
		stars:    make(map[string]*Star),
		planets:  make(map[string]*Planet),
		colonies: make(map[string]*Colony),
		ships:    make(map[string]*Ship),
	}
	cluster.admins["admin"] = true

	polity := polity()
	polity.id = "usagi"
	cluster.polities[polity.id] = polity
	polity.name = polity.id

	// system
	system := mksystem(1, 1, 1)
	system.id = "mizugame"
	cluster.systems[system.id] = system
	polity.home.system = system

	// star
	star := mkstar(system)
	star.id = "shikoku"
	cluster.stars[star.id] = star
	polity.home.star = star

	// fifth orbit planet
	planet := mkplanet(mkorbit(star, 4), TERRESTRIAL)
	planet.id = "suisei"
	planet.name = planet.id
	planet.habitability = 25 // in tens of millions
	cluster.planets[planet.id] = planet
	polity.home.planet = planet

	for j, kind := range []ResourceKind{RFUEL, RGOLD, RMETAL, RNONMETAL} {
		resource := mkresource(kind, true)
		resource.id = fmt.Sprintf("%s-%s-%02d", planet.id, kind, j+1)
		planet.deposits = append(planet.deposits, resource)
	}

	// open colony on planet in fifth orbit
	openColony := mkcolony(polity, nil, planet, OPEN)
	openColony.id = "sanuki"
	cluster.colonies[openColony.id] = openColony
	polity.home.colony = openColony

	openColony.units = append(openColony.units, Unit{Kind: FARM, Assembled: true, TechLevel: 1, Quantity: 500_000})
	openColony.units = append(openColony.units, Unit{Kind: POWER, Assembled: true, TechLevel: 1, Quantity: 1_000_000})
	openColony.units = append(openColony.units, Unit{Kind: MINE, Assembled: true, TechLevel: 1, Quantity: 250_000})
	openColony.units = append(openColony.units, Unit{Kind: MINE, Assembled: true, TechLevel: 1, Quantity: 250_000})
	openColony.units = append(openColony.units, Unit{Kind: MINE, Assembled: true, TechLevel: 1, Quantity: 250_000})
	openColony.units = append(openColony.units, Unit{Kind: MINE, Assembled: true, TechLevel: 1, Quantity: 250_000})

	// tenth orbit colony
	orbit := mkorbit(star, 9)
	orbitingColony := mkcolony(polity, orbit, nil, ENCLOSED)
	orbitingColony.id = "tosa"
	cluster.colonies[orbitingColony.id] = orbitingColony

	for j, kind := range []ResourceKind{RFUEL, RGOLD, RMETAL, RNONMETAL} {
		resource := mkresource(kind, true)
		resource.id = fmt.Sprintf("%s-rsrc-%02d", orbitingColony.id, j+1)
		orbit.deposits = append(orbit.deposits, resource)
	}

	return cluster
}
