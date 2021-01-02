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
	"log"
)

func (st *State) ExecuteOrders(debug bool) []error {
	var errs []error
	for _, err := range st.gameDataCleanupStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.combatOrdersStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.permissionOrdersStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.disassemblyStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.setupStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.transferStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.draftOrdersStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.assemblyStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.buildChangeStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.surveysAndProbesStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.payChangeStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.namingOrdersStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.shipTravelStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.probeStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.giveStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.productionStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.produceOutputStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.sendOutputStage(debug) {
		errs = append(errs, err)
	}
	return errs
}

// Assembly Stage
// = Order Processing Segment
// == Assemble
// == Expend (Research Points only)
// === Expending Research Points to Advance Technology Level (TL)
// == Expend (using Prototypes)
// == PROTOTYPES
// == Factory Group Change
// == Build Change
// == Mine Change
// == Shut Down
// == Start Up
// = Non Prototype TL Increases Segment
// == Expend Research Points only from the Committed Research Buffer
func (st *State) assemblyStage(debug bool) []error {
	stageName := "assembly"
	var errs []error
	for i, order := range st.orders {
		switch {
		case order.Debug != nil:
			debug = order.Debug.On
			if debug {
				log.Printf("[stage:%s] %4d debug %v\n", stageName, i, *order.Debug)
			}
		}
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

// Build Change Stage
func (st *State) buildChangeStage(debug bool) []error {
	stageName := "buildChange"
	var errs []error
	for i, order := range st.orders {
		switch {
		case order.Debug != nil:
			debug = order.Debug.On
			if debug {
				log.Printf("[stage:%s] %4d debug %v\n", stageName, i, *order.Debug)
			}
		}
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

// . Cycles Through Colonies:
//   .. Sums and reports Professionals used to pilot transports.
//   .. Collects data for surveys.
//   .. Totals automation capacity and life support capacity.
//   .. Does production in the following order:
//      ... Power Production
//      ... Mine Production
//      ... Farm Production
//      ... Laboratory Production
//      ... Factory Production
//   .. Food Consumption
//   .. Consumer Goods Consumption (includes ships calling this colony home port)
//   .. Rebel Actions
//   .. Population Changes (Births, Deaths, Graduations & Retirements)
//   .. Statistics updates
func (st *State) colonyProductionStage(debug bool) []error {
	stageName := "colonyProduction"
	var errs []error
	for _, colony := range st.colonies {
		log.Printf("[stage:%s] colony %q\n", stageName, colony.name)
		var totalPopulation int
		totalPopulation += colony.units.population.construction
		totalPopulation += colony.units.population.professionals
		totalPopulation += colony.units.population.soldiers
		totalPopulation += colony.units.population.spies
		totalPopulation += colony.units.population.trainees
		totalPopulation += colony.units.population.unskilled
		totalPopulation += colony.units.population.others

		// farm production
		var unitsProduced, unitsStored int
		for _, farm := range colony.units.farms {
			unitsProduced += farm.Produce()
		}

		// calculate food needed
		unitsNeeded := totalPopulation

		// consume from production before taking from storage
		if unitsProduced >= unitsNeeded {
			unitsProduced, unitsNeeded = unitsProduced-unitsNeeded, 0
		} else {
			unitsProduced, unitsNeeded = 0, unitsNeeded-unitsProduced
			if unitsStored <= unitsNeeded {
				unitsStored, unitsNeeded = 0, unitsNeeded-unitsStored
			} else {
				unitsStored, unitsNeeded = unitsStored-unitsNeeded, 0
			}
		}

		if unitsNeeded != 0 {
			// starvation
		}
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

// Combat Orders Stage
// = Prefire Segment
// == Dodge
// == Accept
// == Auto Return Fire
// == Close Proximity Targeting
// = Pre-Maneuver Fire Segment
// == Pre-maneuver Energy Weapon Fire
// === Bombardment Target Categories
// == Pre-maneuver Missile Fire
// === Bombardment Target Categories
// == Allocate Damage
// = Maneuver Segment
// == Undock
// == Run
// == Tactical Maneuver
// == Close
// == Dock
// == Undocking
// == Allocate Damage
// = Post-Maneuver Fire Segment
// == After-maneuver Energy Weapon Fire
// === Bombardment Target Categories
// == After-maneuver Missile Fire
// === Bombardment Target Categories
// == Allocate Damage
// = Ground Combat Segment
// == Withdraw
// == Defensive Support
// == Invasion
// == Offensive Support
// == Cycle Ground Combat
// === COMBAT EXECUTION SEQUENCE
// === PRE-MANEUVER FIRE
// === SHIP MANEUVERS
// === POST-MANEUVER FIRE
// === TROOP MOVEMENT
// === ANTI-INVASION ENERGY WEAPON FIRE
// === MILITIA
// === SURRENDER CHECK
// === INVASION CASUALTIES
// === CALCULATING COMBAT FACTORS
// === DETERMINING UNIT LOSSES IN COMBAT
// === MISSION COMPLETION
// === Bombardment Using Ranged Weaponry
// === Combat Using Military Units
// === Damage Allocation
func (st *State) combatOrdersStage(debug bool) []error {
	stageName := "combatOrders"
	var errs []error
	for i, order := range st.orders {
		switch {
		case order.Debug != nil:
			debug = order.Debug.On
			if debug {
				log.Printf("[stage:%s] %4d debug %v\n", stageName, i, *order.Debug)
			}
		}
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

func (st *State) combineFactoryGroupStage(debug bool) []error {
	stageName := "combineFactoryGroup"
	var errs []error
	for i, order := range st.orders {
		switch {
		case order.Debug != nil:
			debug = order.Debug.On
			if debug {
				log.Printf("[stage:%s] %4d debug %v\n", stageName, i, *order.Debug)
			}
		}
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

func (st *State) disassembleStage(debug bool) []error {
	stageName := "disassemble"
	var errs []error
	for i, order := range st.orders {
		switch {
		case order.Debug != nil:
			debug = order.Debug.On
			if debug {
				log.Printf("[stage:%s] %4d debug %v\n", stageName, i, *order.Debug)
			}
		}
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

// Disassembly Stage
func (st *State) disassemblyStage(debug bool) []error {
	stageName := "disassembly"
	var errs []error
	for _, err := range st.disassembleStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.scrapStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.junkStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.mergeStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.combineFactoryGroupStage(debug) {
		errs = append(errs, err)
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

func (st *State) disbandStage(debug bool) []error {
	stageName := "disband"
	var errs []error
	for i, order := range st.orders {
		switch {
		case order.Debug != nil:
			debug = order.Debug.On
			if debug {
				log.Printf("[stage:%s] %4d debug %v\n", stageName, i, *order.Debug)
			}
		}
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

func (st *State) draftStage(debug bool) []error {
	stageName := "draft"
	var errs []error
	for i, order := range st.orders {
		switch {
		case order.Debug != nil:
			debug = order.Debug.On
			if debug {
				log.Printf("[stage:%s] %4d debug %v\n", stageName, i, *order.Debug)
			}
		}
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

// Draft Orders Stage
func (st *State) draftOrdersStage(debug bool) []error {
	stageName := "draftOrders"
	var errs []error
	for _, err := range st.draftStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.disbandStage(debug) {
		errs = append(errs, err)
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

// Game Data Cleanup Stage
func (st *State) gameDataCleanupStage(debug bool) []error {
	stageName := "gameDataCleanup"
	var errs []error
	for i, order := range st.orders {
		switch {
		case order.Debug != nil:
			debug = order.Debug.On
			if debug {
				log.Printf("[stage:%s] %4d debug %v\n", stageName, i, *order.Debug)
			}
		}
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

// Give Stage
func (st *State) giveStage(debug bool) []error {
	stageName := "give"
	var errs []error
	for i, order := range st.orders {
		switch {
		case order.Debug != nil:
			debug = order.Debug.On
			if debug {
				log.Printf("[stage:%s] %4d debug %v\n", stageName, i, *order.Debug)
			}
		case order.Give != nil:
			if debug {
				log.Printf("[stage:%s] %4d give %v\n", stageName, i, *order.Give)
			}
			if err := st.Give(order.issuedBy, order.Give.AssetID, order.Give.TargetID); err != nil {
				errs = append(errs, err)
			}
		}
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

func (st *State) jumpStage(debug bool) []error {
	stageName := "jump"
	var errs []error
	for i, order := range st.orders {
		switch {
		case order.Debug != nil:
			debug = order.Debug.On
			if debug {
				log.Printf("[stage:%s] %4d debug %v\n", stageName, i, *order.Debug)
			}
		}
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

func (st *State) junkStage(debug bool) []error {
	stageName := "junk"
	var errs []error
	for i, order := range st.orders {
		switch {
		case order.Debug != nil:
			debug = order.Debug.On
			if debug {
				log.Printf("[stage:%s] %4d debug %v\n", stageName, i, *order.Debug)
			}
		}
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

func (st *State) loadCargoStage(debug bool) []error {
	stageName := "loadCargo"
	var errs []error
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

func (st *State) mergeStage(debug bool) []error {
	stageName := "merge"
	var errs []error
	for i, order := range st.orders {
		switch {
		case order.Debug != nil:
			debug = order.Debug.On
			if debug {
				log.Printf("[stage:%s] %4d debug %v\n", stageName, i, *order.Debug)
			}
		}
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

func (st *State) moveStage(debug bool) []error {
	stageName := "move"
	var errs []error
	for i, order := range st.orders {
		switch {
		case order.Debug != nil:
			debug = order.Debug.On
			if debug {
				log.Printf("[stage:%s] %4d debug %v\n", stageName, i, *order.Debug)
			}
		}
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

// Naming Orders Stage
// = Name
// = Note
// = Control Planet
// = Un-control Planet
// = Message
func (st *State) namingOrdersStage(debug bool) []error {
	stageName := "namingOrders"
	var errs []error
	for i, order := range st.orders {
		switch {
		case order.Debug != nil:
			debug = order.Debug.On
			if debug {
				log.Printf("[stage:%s] %4d debug %v\n", stageName, i, *order.Debug)
			}
		}
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

func (st *State) payStage(debug bool) []error {
	stageName := "pay"
	var errs []error
	for i, order := range st.orders {
		switch {
		case order.Debug != nil:
			debug = order.Debug.On
			if debug {
				log.Printf("[stage:%s] %4d debug %v\n", stageName, i, *order.Debug)
			}
		}
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

// Pay Change Stage
func (st *State) payChangeStage(debug bool) []error {
	stageName := "payChange"
	var errs []error
	for _, err := range st.payStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.rationStage(debug) {
		errs = append(errs, err)
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

// Permission Orders Stage
// = Permission to Colonize
// = Home Port Change
// = Diplomacy
func (st *State) permissionOrdersStage(debug bool) []error {
	stageName := "permissionOrders"
	var errs []error
	for i, order := range st.orders {
		switch {
		case order.Debug != nil:
			debug = order.Debug.On
			if debug {
				log.Printf("[stage:%s] %4d debug %v\n", stageName, i, *order.Debug)
			}
		}
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

func (st *State) pickupStage(debug bool) []error {
	stageName := "pickup"
	var errs []error
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

// Probe Stage
// = Probes
// == Orbit Probe
// == Star System Probe
// == S/C or "Intensive" Probe
func (st *State) probeStage(debug bool) []error {
	stageName := "probe"
	var errs []error
	for i, order := range st.orders {
		switch {
		case order.Debug != nil:
			debug = order.Debug.On
			if debug {
				log.Printf("[stage:%s] %4d debug %v\n", stageName, i, *order.Debug)
			}
		}
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

// Produce Output Stage
func (st *State) produceOutputStage(debug bool) []error {
	stageName := "produceOutput"
	var errs []error
	for i, order := range st.orders {
		switch {
		case order.Debug != nil:
			debug = order.Debug.On
			if debug {
				log.Printf("[stage:%s] %4d debug %v\n", stageName, i, *order.Debug)
			}
		}
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

// Production Stage
func (st *State) productionStage(debug bool) []error {
	stageName := "production"
	var errs []error
	for _, err := range st.colonyProductionStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.shipProductionStage(debug) {
		errs = append(errs, err)
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

func (st *State) rationStage(debug bool) []error {
	stageName := "ration"
	var errs []error
	for i, order := range st.orders {
		switch {
		case order.Debug != nil:
			debug = order.Debug.On
			if debug {
				log.Printf("[stage:%s] %4d debug %v\n", stageName, i, *order.Debug)
			}
		}
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

func (st *State) scrapStage(debug bool) []error {
	stageName := "scrap"
	var errs []error
	for i, order := range st.orders {
		switch {
		case order.Debug != nil:
			debug = order.Debug.On
			if debug {
				log.Printf("[stage:%s] %4d debug %v\n", stageName, i, *order.Debug)
			}
		}
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

// Send Output Stage
func (st *State) sendOutputStage(debug bool) []error {
	stageName := "sendOutput"
	var errs []error
	for i, order := range st.orders {
		switch {
		case order.Debug != nil:
			debug = order.Debug.On
			if debug {
				log.Printf("[stage:%s] %4d debug %v\n", stageName, i, *order.Debug)
			}
		}
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

// Setup Stage
// = Define Cargo Hold
// = Set Up
// == SET UP ORDERS
// === SET UP DESCRIPTION
// === SET UP RESTRICTIONS
// === RESTRICTIONS ON THE TYPE OF COLONY
// === PERMISSION TO COLONIZE
// === FAILURES TO SET UP
// === PERSONNEL
// === CONSTRUCTORS AND SET UPS
// === TRANSPORTS AND SET UPS
// === ASSEMBLIES AND SET UPS
// = Add On
// == ADD-ON DETAILS
// == DESCRIPTION
// == ADD-ON RESTRICTIONS
// == ADD-ON FAILURES
// == HOW ADD-ON ORDERS WORK
// == CONSTRUCTORS AND ADD-ON
// == TRANSPORTS AND ADD-ONS
// == ASSEMBLIES AND ADD-ONS
// == ADD-ON LOCATIONS
func (st *State) setupStage(debug bool) []error {
	stageName := "setup"
	var errs []error
	for i, order := range st.orders {
		switch {
		case order.Debug != nil:
			debug = order.Debug.On
			if debug {
				log.Printf("[stage:%s] %4d debug %v\n", stageName, i, *order.Debug)
			}
		}
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

// shipProductionStage calculates production and consumption for all ships.
//
// For each ship:
//   1. Sums and reports Professionals used to pilot transports.
//   2. Totals automation capacity and life support capacity.
//   3. Farm Production
//   4. Food Consumption
//   5. Rebel Actions
//   6. Population Changes (Deaths Graduations Retirements)
//   7. Statistics updates
func (st *State) shipProductionStage(debug bool) []error {
	stageName := "shipProduction"
	var errs []error
	for _, ship := range st.ships {
		log.Printf("[stage:%s] ship %q\n", stageName, ship.name)
		var totalPopulation int
		totalPopulation += ship.units.population.construction
		totalPopulation += ship.units.population.professionals
		totalPopulation += ship.units.population.soldiers
		totalPopulation += ship.units.population.spies
		totalPopulation += ship.units.population.trainees
		totalPopulation += ship.units.population.unskilled
		totalPopulation += ship.units.population.others

		// farm production
		var unitsProduced, unitsStored int
		for _, farm := range ship.units.farms {
			unitsProduced += farm.Produce()
		}

		// calculate food needed
		unitsNeeded, unitsConsumed := totalPopulation, 0
		unitsRationed := int(float64(unitsNeeded) * ship.ration)
		if unitsRationed != unitsNeeded {
			// potential for starvation
		}

		// consume food from production before taking any from storage
		if unitsProduced >= unitsNeeded {
			unitsProduced, unitsNeeded = unitsProduced-unitsNeeded, 0
		} else {
			unitsProduced, unitsNeeded = 0, unitsNeeded-unitsProduced
			if unitsStored <= unitsNeeded {
				unitsStored, unitsNeeded = 0, unitsNeeded-unitsStored
			} else {
				unitsStored, unitsNeeded = unitsStored-unitsNeeded, 0
			}
		}

		if unitsConsumed != 0 {
			// okay
		}
		if unitsNeeded != 0 { // calculate deaths due to starvation
		}
		if unitsProduced != 0 { // move to storage, any excess is wasted.
		}

		// rebel actions
		// population changes
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

// Ship Travel Stage
func (st *State) shipTravelStage(debug bool) []error {
	stageName := "shipTravel"
	var errs []error
	for _, err := range st.jumpStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.moveStage(debug) {
		errs = append(errs, err)
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

func (st *State) shortagesStage(debug bool) []error {
	stageName := "shortages"
	var errs []error
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

// Surveys and Probes Stage
// = S/C Probes Only
// = Survey
// = Launch Robot Probe
// == ORBIT PROBE
// == STAR SYSTEM PROBE
// == S/C PROBE
// == SURVEY
// == RPV TL to distance formulas
func (st *State) surveysAndProbesStage(debug bool) []error {
	stageName := "surveysAndProbes"
	var errs []error
	for i, order := range st.orders {
		switch {
		case order.Debug != nil:
			debug = order.Debug.On
			if debug {
				log.Printf("[stage:%s] %4d debug %v\n", stageName, i, *order.Debug)
			}
		}
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

// Transfer Stage
// = Unload Cargo
// = Transfer
// == TRANSFERS AND PICK-UPS
// == TRANSPORTS
// == TRANSPORT CAPACITY
// == TRANSFERRING OR PICKING-UP TRANSPORTS
// == SHORTAGES
// = Pick Up
// = Load Cargo
func (st *State) transferStage(debug bool) []error {
	stageName := "transfer"
	var errs []error
	for _, err := range st.unloadCargoStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.transferUnitsStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.pickupStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.loadCargoStage(debug) {
		errs = append(errs, err)
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

func (st *State) transferAndPickupStage(debug bool) []error {
	stageName := "transferAndPickup"
	var errs []error
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

func (st *State) transferUnitsStage(debug bool) []error {
	stageName := "transferUnits"
	var errs []error
	for _, err := range st.transferAndPickupStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.transportsStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.transportCapacityStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.transferOrPickupTransportsStage(debug) {
		errs = append(errs, err)
	}
	for _, err := range st.shortagesStage(debug) {
		errs = append(errs, err)
	}
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

func (st *State) transferOrPickupTransportsStage(debug bool) []error {
	stageName := "transferOrPickupTransports"
	var errs []error
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

func (st *State) transportCapacityStage(debug bool) []error {
	stageName := "transportCapacity"
	var errs []error
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

func (st *State) transportsStage(debug bool) []error {
	stageName := "transports"
	var errs []error
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}

func (st *State) unloadCargoStage(debug bool) []error {
	stageName := "unloadCargo"
	var errs []error
	return append(errs, fmt.Errorf("%s: %w", stageName, ERRNOTIMPLEMENTED))
}
