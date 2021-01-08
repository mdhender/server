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

package population

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mdhender/server/pkg/utils"
)

func New(options ...Option) Population {
	var p Population
	for _, option := range options {
		p = option(p)
	}
	return p
}

type Option func(population Population) Population

func NewHomeColony(orbiting bool) Population {
	var p Population
	if orbiting {
		p.construction = 10_000
		p.professionals = 100_000
		p.soldiers = 150_000
		p.unskilled = 370_000
		p.others = 350_000
	} else {
		p.construction = 20_000
		p.professionals = 2_000_000
		p.soldiers = 2_500_000
		p.unskilled = 6_000_000
		p.others = 5_900_000
	}
	p.total = p.construction + p.professionals + p.soldiers + p.spies + p.trainees + p.unskilled + p.others
	return p
}

func Set(k Kind, n int) func(p Population) Population {
	if n < 0 {
		n = 0
	}
	return func(p Population) Population {
		switch k {
		case CONSTRUCTION:
			p.construction = n
		case PROFESSIONALS:
			p.professionals = n
		case SOLDIERS:
			p.soldiers = n
		case SPIES:
			p.spies = n
		case TRAINEES:
			p.trainees = n
		case UNSKILLED:
			p.unskilled = n
		default: // OTHERS
			p.others = n
		}
		p.total += n
		return p
	}
}

// Count returns the current number of units for the kind requested.
func (p Population) Count(k Kind) int {
	switch k {
	case CONSTRUCTION:
		return p.construction
	case PROFESSIONALS:
		return p.professionals
	case SOLDIERS:
		return p.soldiers
	case SPIES:
		return p.spies
	case TRAINEES:
		return p.trainees
	case UNSKILLED:
		return p.unskilled
	default: // OTHERS
		return p.others
	}
}

// TotalCount returns the total number of units in the population
func (p Population) TotalCount() int {
	return p.total
}

// FoodNeededPerQuarter returns the number of food units needed to
// fully feed a population for one game turn as well as the minimum
// amount needed to prevent starvation.
//
// The full amount is 1.00 food units per population per turn.
// The minimum amount needed to avoid starvation is is 0.25 food units per population per turn.
func (p Population) FoodNeededPerTurn() (min, full int) {
	min = p.total / 4
	if min*4 != p.total {
		min++
	}
	return min, p.total
}

// FoodStockpileGoal returns the number of food units that the
// population wants to always have on hand. It's equal to a four
// turns of full rations. Having less than this amount tends to
// increase the discontent of the population.
func (p Population) FoodStockpileGoal() int {
	return p.total
}

// Population is the number and type of population within a ship or colony.
type Population struct {
	construction  int
	professionals int
	soldiers      int
	spies         int
	trainees      int
	unskilled     int
	others        int
	total         int
}

// Kind is the type of population unit.
// It controls what actions the unit may perform.
type Kind int

const (
	OTHERS Kind = iota
	CONSTRUCTION
	PROFESSIONALS
	SOLDIERS
	SPIES
	TRAINEES
	UNSKILLED
)

func (p Population) String() string {
	return fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%s",
		utils.Commas(p.construction),
		utils.Commas(p.professionals),
		utils.Commas(p.soldiers),
		utils.Commas(p.spies),
		utils.Commas(p.trainees),
		utils.Commas(p.unskilled),
		utils.Commas(p.others),
		utils.Commas(p.total),
	)
}

func (k Kind) String() string {
	switch k {
	case CONSTRUCTION:
		return "construction"
	case PROFESSIONALS:
		return "professionals"
	case SOLDIERS:
		return "soldiers"
	case SPIES:
		return "spies"
	case TRAINEES:
		return "trainees"
	case UNSKILLED:
		return "unskilled"
	default: // OTHERS
		return "others"
	}
}

// MarshalJSON implements the json.Marshaler interface.
func (p Population) MarshalJSON() ([]byte, error) {
	inout := struct {
		Construction  int `json:"construction,omitempty"`
		Professionals int `json:"professionals,omitempty"`
		Soldiers      int `json:"soldiers,omitempty"`
		Spies         int `json:"spies,omitempty"`
		Trainees      int `json:"trainees,omitempty"`
		Unskilled     int `json:"unskilled,omitempty"`
		Others        int `json:"others,omitempty"`
		Total         int `json:"total"`
	}{
		Construction:  p.construction,
		Professionals: p.professionals,
		Soldiers:      p.soldiers,
		Spies:         p.spies,
		Trainees:      p.trainees,
		Unskilled:     p.unskilled,
		Others:        p.others,
		Total:         p.total,
	}
	return json.Marshal(&inout)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// Invalid values in the input are treated as zero values.
// It ignores extra fields in the input?
// The totals value is derived rather than trusting the input.
func (p *Population) UnmarshalJSON(b []byte) error {
	// per the docs, a value of "null" is treated as a no-op.
	// in our case, that means zero values for the population.
	if bytes.Equal(b, []byte("null")) {
		return nil
	}

	var inout struct {
		Construction  int `json:"construction"`
		Professionals int `json:"professionals"`
		Soldiers      int `json:"soldiers"`
		Spies         int `json:"spies"`
		Trainees      int `json:"trainees"`
		Unskilled     int `json:"unskilled"`
		Others        int `json:"others"`
	}
	if err := json.Unmarshal(b, &inout); err != nil {
		return err
	}
	if inout.Construction > 0 {
		p.construction = inout.Construction
		p.total += inout.Construction
	}
	if inout.Professionals > 0 {
		p.professionals = inout.Professionals
		p.total += inout.Professionals
	}
	if inout.Soldiers > 0 {
		p.soldiers = inout.Soldiers
		p.total += inout.Soldiers
	}
	if inout.Spies > 0 {
		p.spies = inout.Spies
		p.total += inout.Spies
	}
	if inout.Trainees > 0 {
		p.trainees = inout.Trainees
		p.total += inout.Trainees
	}
	if inout.Unskilled > 0 {
		p.unskilled = inout.Unskilled
		p.total += inout.Unskilled
	}
	if inout.Others > 0 {
		p.others = inout.Others
		p.total += inout.Others
	}
	return nil
}
