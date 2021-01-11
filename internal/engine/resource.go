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

import "github.com/google/uuid"

func mkresource(kind ResourceKind, unlimited bool) *Resource {
	resource := &Resource{
		id:            uuid.New().String(),
		kind:          kind,
		unlimited:     unlimited,
		initialAmount: 55 * 1_000_000_000,
		yieldPct:      0.90,
	}

	if resource.kind == RGOLD {
		resource.yieldPct /= 10
		resource.initialAmount /= 10
	}

	resource.amountRemaining = resource.initialAmount

	if resource.unlimited {
		resource.yieldPct /= 3
	}

	return resource
}

// Resource is any resource that can be mined.
type Resource struct {
	id              string
	kind            ResourceKind
	unlimited       bool
	initialAmount   int
	amountRemaining int
	yieldPct        float64
}
