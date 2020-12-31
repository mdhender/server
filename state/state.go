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

package state

type State struct {
}

type System struct {
	ID     string
	Name   string
	Coords Coords
	Stars  []*Star
}

type Coords struct {
	X int
	Y int
	Z int
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

// Resource is any resource that can be mined.
type Resource struct {
	ID              string
	Type            ResourceType
	YieldPct        float64
	Unlimited       bool
	InitialAmount   int64
	AmountRemaining int64
}

type Race struct {
	Name string
	Home struct {
		System *System
		World  *Planet
	}
}

type AutomationUnit struct{}
type Colony struct{}
type EngineUnit struct{}
type FactoryUnit struct{}
type FarmUnit struct{}
type MineUnit struct{}
type MissileUnit struct{}
type AntiMissileUnit struct{}
type PopulationUnit struct{}
type RobotUnit struct{}
type Ship struct{}
type StructuralUnit struct{}
type TransportUnit struct{}

// enums
type PlanetType int
type ResourceType int
