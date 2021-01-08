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
	"github.com/matryer/is"
	"testing"
)

func Test_Make(t *testing.T) {
	is := is.New(t)

	st, admin := Make()
	is.True(admin != "")
	is.True(st.admins[admin])
}

func Test_CreateAdmin(t *testing.T) {
	is := is.New(t)

	st, admin := Make()
	order := Order{CreateAdmin: &CreateAdmin{}}
	order.Stamp(admin)
	is.True(order.issuedBy == admin)

	var orders Orders
	orders = append(orders, &order)
	is.True(len(orders) == 1)

	id, errs := st.CreateAdmin(admin, "mdhender")
	is.True(len(errs) == 0)
	is.True(id == "mdhender")
	is.True(st.admins[id])
}

func Test_CreatePolity(t *testing.T) {
	is := is.New(t)

	st, admin := Make()
	id, errs := st.CreatePolity(admin, "")
	is.True(len(errs) == 0)
	is.True(id != "")
	p := st.LookupPolity(id)
	is.True(p != nil)
	is.True(p.ID() == id)
}
