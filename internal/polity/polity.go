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

package polity

import "github.com/google/uuid"

type Polity struct {
	id string
	serves string // set only if the polity is a viceroy
}

func New(id string) *Polity {
	if id == "" {
		id = uuid.New().String()
	}

	return &Polity{
		id: id,
	}
}

func (p *Polity) ID() string {
	return p.id
}

func (p *Polity) Serves(id string) bool {
	return id != "" && id == p.serves
}