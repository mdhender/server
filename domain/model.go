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

package domain

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
	ID     string
	Planet *Planet
}

type Planet struct {
	ID           string
	Type         PlanetType
	Habitability int // range from 0 to 25
	Deposits     []*Resource
}

type PlanetType int

// Resource is any resource that can be mined.
type Resource struct {
	ID              string
	Type            ResourceType
	YieldPct        float64
	Unlimited       bool
	InitialAmount   int64
	AmountRemaining int64
}

type ResourceType int

type Race struct {
	Name string
	Home struct {
		System *System
		World  *Planet
	}
}
