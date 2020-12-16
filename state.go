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
	"log"
)

type State struct{}

func (e *State) PostOrders(orders Orders) error {
	// do something with the orders
	var debug bool
	for i, order := range orders {
		switch {
		case order.Debug != nil:
			debug = order.Debug.On
			if debug {
				log.Printf("[orders] %4d debug %v\n", i, *order.Debug)
			}
		case order.DefensiveSupport != nil:
			if debug {
				log.Printf("[orders] %4d defensiveSupport %v\n", i, *order.DefensiveSupport)
			}
		case order.Dock != nil:
			if debug {
				log.Printf("[orders] %4d dock %v\n", i, *order.Dock)
			}
		case order.Run != nil:
			if debug {
				log.Printf("[orders] %4d run %v\n", i, *order.Run)
			}
		case order.Undock != nil:
			if debug {
				log.Printf("[orders] %4d undock %v\n", i, *order.Undock)
			}
		}
	}
	return nil
}
