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

import (
	"encoding/json"
)

func (p Planet) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID           string      `json:"planet_id"`
		Type         string      `json:"type"`
		Habitability int         `json:"habitability"`
		Deposits     []*Resource `json:"deposits"`
	}{
		ID:           p.ID,
		Type:         p.Type.String(),
		Habitability: p.Habitability,
		Deposits:     p.Deposits,
	})
}

func (r *Race) MarshalJSON() ([]byte, error) {
	var data struct {
		Name string `json:"name"`
		Home struct {
			System string `json:"system"`
			World  string `json:"world"`
		} `json:"home"`
	}
	data.Name = r.Name
	data.Home.System = r.Home.System.Name
	data.Home.World = r.Home.World.ID
	return json.Marshal(&data)
}

func (nr Resource) MarshalJSON() ([]byte, error) {
	amtRemaining := nr.AmountRemaining
	if nr.Unlimited {
		amtRemaining = 99_999_999
	}
	return json.Marshal(&struct {
		ID              string `json:"deposit_id"`
		Type            string `json:"type"`
		Yield           int    `json:"yield"`
		AmountRemaining int64  `json:"remaining"`
	}{
		ID:              nr.ID,
		Type:            nr.Type.String(),
		Yield:           int(nr.YieldPct * 100),
		AmountRemaining: amtRemaining,
	})
}

