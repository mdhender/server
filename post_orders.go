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

import "log"

func (st *State) PostOrders(orders Orders) error {
	// do something with the orders
	var debug bool
	var errs []error
	for i, order := range orders {
		var orderedBy string // todo: set to the player that issued the order

		switch {
		case order.Debug != nil:
			debug = order.Debug.On
			if debug {
				log.Printf("[orders] %4d debug %v\n", i, *order.Debug)
			}
		case order.Accept != nil:
			if debug {
				log.Printf("[orders] %4d accept %v\n", *order.Accept)
			}
			if err := st.Accept(orderedBy, order.Accept.AssetID); err != nil {
				errs = append(errs, err)
			}
		case order.DefensiveSupport != nil:
			if debug {
				log.Printf("[orders] %4d defensiveSupport %v\n", i, *order.DefensiveSupport)
			}
		case order.Dock != nil:
			if debug {
				log.Printf("[orders] %4d dock %v\n", i, *order.Dock)
			}
		case order.Give != nil:
			if debug {
				log.Printf("[orders] %4d give %v\n", i, *order.Give)
			}
			if err := st.Give(orderedBy, order.Give.AssetID, order.Give.TargetID); err != nil {
				errs = append(errs, err)
			}
		case order.HomePortChange != nil:
			if debug {
				log.Printf("[orders] %4d homePortChange %v\n", i, *order.HomePortChange)
			}
			if err := st.HomePortChange(orderedBy, order.HomePortChange.ShipID, order.HomePortChange.ColonyID); err != nil {
				errs = append(errs, err)
			}
		case order.Junk != nil:
			if debug {
				log.Printf("[orders] %4d junk %v\n", i, *order.Junk)
			}
			if err := st.Junk(orderedBy, order.Junk.ActorID, order.Junk.AssetID); err != nil {
				errs = append(errs, err)
			}
		case order.Name != nil:
			if debug {
				log.Printf("[orders] %4d name %v\n", i, *order.Name)
			}
			if err := st.Name(orderedBy, order.Name.EntityID, order.Name.Type, order.Name.Name); err != nil {
				errs = append(errs, err)
			}
		case order.Note != nil:
			if debug {
				log.Printf("[orders] %4d note %v\n", i, *order.Note)
			}
			if note, err := NewText(order.Note.Text); err != nil {
				errs = append(errs, err)
			} else {
				log.Printf("[todo] State.PostOrders: considering untaining note text\n")
				if err = st.Note(orderedBy, order.Note.TargetID, note); err != nil {
					errs = append(errs, err)
				}
			}
		case order.PermissionToColonize != nil:
			if debug {
				log.Printf("[orders] %4d permissionToColonize %v\n", i, *order.PermissionToColonize)
			}
			if err := st.PermissionToColonize(orderedBy, order.PermissionToColonize.PlanetID, order.PermissionToColonize.ShipID); err != nil {
				errs = append(errs, err)
			}
		case order.Run != nil:
			if debug {
				log.Printf("[orders] %4d run %v\n", i, *order.Run)
			}
		case order.Scrap != nil:
			if debug {
				log.Printf("[orders] %4d scrap %v\n", i, *order.Scrap)
			}
			if err := st.Scrap(orderedBy, order.Scrap.ActorID, order.Scrap.Item, order.Scrap.TechLevel, order.Scrap.Quantity); err != nil {
				errs = append(errs, err)
			}
		case order.Undock != nil:
			if debug {
				log.Printf("[orders] %4d undock %v\n", i, *order.Undock)
			}
		}
	}
	if len(errs) != 0 {
		for _, err := range errs {
			log.Printf("[orders] %+v\n", err)
		}
		return errs[0]
	}
	return nil
}