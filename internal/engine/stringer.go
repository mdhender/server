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
	"fmt"
	"github.com/mdhender/server/pkg/utils"
	"strings"
)

func (st *State) String() string {
	w := &strings.Builder{}
	w.Grow(10 * 1024)
	_, _ = fmt.Fprintf(w, "(state (turn %d)\n", st.turn)
	for _, polity := range st.polities {
		_, _ = fmt.Fprintf(w, "  (polity (id %q)\n", polity.id)
		_, _ = fmt.Fprintf(w, "    (name %q)\n", polity.name)
		_, _ = fmt.Fprintf(w, "    (home (system %q)\n", polity.home.system.id)
		_, _ = fmt.Fprintf(w, "          (world  %q)\n", polity.home.world)
		_, _ = fmt.Fprintf(w, "          (colony %q))\n", polity.home.colony)
		for _, c := range polity.controls.colonies {
			_, _ = fmt.Fprintf(w, "    (colony (id %q)\n", c.id)
			_, _ = fmt.Fprintf(w, "      (hull-number %q)\n", c.number)
			_, _ = fmt.Fprintf(w, "      (kind        %s)\n", c.kind)
			_, _ = fmt.Fprintf(w, "      (ration      %7s)\n", utils.Percentage(c.ration))
			//_,_=fmt.Fprintf(w, "    (population\n")
			//_,_=fmt.Fprintf(w, "      (%-13s %13s)\n", "construction", utils.Commas(c.population.Count(population.CONSTRUCTION)))
			//_,_=fmt.Fprintf(w, "      (%-13s %13s)\n", "professionals", utils.Commas(c.population.Count(population.PROFESSIONALS)))
			//_,_=fmt.Fprintf(w, "      (%-13s %13s)\n", "soldiers", utils.Commas(c.population.Count(population.SOLDIERS)))
			//_,_=fmt.Fprintf(w, "      (%-13s %13s)\n", "spies", utils.Commas(c.population.Count(population.SPIES)))
			//_,_=fmt.Fprintf(w, "      (%-13s %13s)\n", "trainees", utils.Commas(c.population.Count(population.TRAINEES)))
			//_,_=fmt.Fprintf(w, "      (%-13s %13s)\n", "unskilled", utils.Commas(c.population.Count(population.UNSKILLED)))
			//_,_=fmt.Fprintf(w, "      (%-13s %13s)\n", "others", utils.Commas(c.population.Count(population.OTHERS)))
			//_,_=fmt.Fprintf(w, "      (%-13s %13s)\n", "total", utils.Commas(c.population.TotalCount()))
			//fmin, fmax := c.population.FoodNeededPerTurn()
			//_,_=fmt.Fprintf(w, "      (food (min %s) (full %s) (want %s)))\n", utils.Commas(fmin), utils.Commas(fmax), utils.Commas(c.population.FoodStockpileGoal()))
			//_,_=fmt.Fprintf(w, "    (factories)\n")
			//_,_=fmt.Fprintf(w, "    (farms)\n")
			//_,_=fmt.Fprintf(w, "    (mines)\n")
			//_,_=fmt.Fprintf(w, "    (power)\n")
			//_,_=fmt.Fprintf(w, "    (storage\n")
			//_,_=fmt.Fprintf(w, "      (%-13s %13s)\n", "food", utils.Commas(c.storage.food))
			//_,_=fmt.Fprintf(w, "      (%-13s %13s)))\n", "foodGoal", utils.Commas(c.foodStockpileGoal))
			_, _ = fmt.Fprintf(w, "    ) ;; colony %s\n", c.id)
		}
		_, _ = fmt.Fprintf(w, "  ) ;; polity %s\n", polity.id)
	}
	_, _ = fmt.Fprintln(w, "")
	for _, planet := range st.planets {
		_, _ = fmt.Fprintf(w, "  (planet (id %q)\n", planet.id)
		_, _ = fmt.Fprintf(w, "      (name       %q)\n", planet.Name)
		for _, c := range planet.colonies {
			_, _ = fmt.Fprintf(w, "      (colony (id %q))\n", c.id)
		}
		_, _ = fmt.Fprintf(w, "  ) ;; planet %s\n", planet.id)
	}
	_, _ = fmt.Fprintln(w, "")
	for _, c := range st.colonies {
		_, _ = fmt.Fprintf(w, "  (colony (id %q)\n", c.id)
		_, _ = fmt.Fprintf(w, "      (kind        %s)\n", c.kind)
		_, _ = fmt.Fprintf(w, "      (hull-number %q)\n", c.number)
		_, _ = fmt.Fprintf(w, "      (name        %q)\n", c.name)
		//_, _ = fmt.Fprintf(w, "      (created-by  %q)\n", c.CreatedBy())
		//_, _ = fmt.Fprintf(w, "      (owned-by    %q)\n", c.OwnedBy)
		_, _ = fmt.Fprintf(w, "      (ration      %7s)\n", utils.Percentage(c.ration))
		//if len(c.units) != 0 {
		//	_, _ = fmt.Fprintf(w, "      (units\n")
		//	for _, u := range c.Units {
		//		_, _ = fmt.Fprintf(w, "        %s\n", u.Sexpr())
		//	}
		//	_, _ = fmt.Fprintf(w, "      ) ;; units\n")
		//}
		_, _ = fmt.Fprintf(w, "  ) ;; colony %s\n", c.id)
	}
	_, _ = fmt.Fprintf(w, ") ;; turn %d\n", st.turn)
	return w.String()
}
