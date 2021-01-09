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

package main

import (
	"errors"
	"fmt"
	"github.com/mdhender/server/internal/engine"
	"github.com/mdhender/server/internal/storage/memory"
	"log"
	"net/http"
)

func run(cfg *config) error {
	rc := routeConfig{
		gameFileSavePath: cfg.Games.FileSavePath,
	}

	ds, err := memory.New()
	if err != nil {
		return err
	}
	if cfg.MockData {
		ds.MockData()
	}
	rc.services.adding = ds
	rc.services.listing = ds
	rc.services.updating = ds

	var options []func(*server) error
	options = append(options, setSalt(cfg.Server.Salt))

	srv, err := newServer(cfg, options...)
	if err != nil {
		return err
	}
	srv.Handler = CorsHandler(routes(srv, rc))

	admin := cfg.Setup.DefaultAdmin
	st, err := engine.NewState(admin)
	if err != nil {
		return fmt.Errorf("engine: %w", err)
	}
	log.Printf("[run] state created with default admin of %q\n", admin)

	var os engine.Orders
	os = append(os, (&engine.Order{CreateAdmin: &engine.CreateAdmin{ID: "mdhender"}}).Stamp(admin))

	var scarcity bool
	for _, input := range []*struct {
		polity  string
		system  string
		x, y, z int
	}{
		{"usagi", "Shikoku", 1, 1, 1},
		{"tomoe", "Kyushu", 2, 2, 2},
	} {
		os = append(os, (&engine.Order{CreateSystem: &engine.CreateSystem{X: input.x, Y: input.y, Z: input.z}}).Stamp(admin))
		os = append(os, (&engine.Order{CreatePolity: &engine.CreatePolity{ID: input.polity, Name: input.polity}}).Stamp(admin))
		log.Printf("[state] polity %q\n", input.polity)

		//system := st.MakeSystem(input.sysname, input.x, input.y, input.z)
		//log.Printf("[state] system %q\n", system.name)
		//polity.home.system = system
		//
		//system.stars[0].orbits[5].planet = st.MakeHomePlanet(polity, scarcity)

		scarcity = !scarcity
	}

	os.Prioritize()
	if st, errs := st.ProcessOrders(os, true); len(errs) != 0 {
		fmt.Printf("errors -----------------------------------------------------\n")
		var counter int
		for _, err := range errs {
			if !errors.Is(err, engine.ERRNOTIMPLEMENTED) {
				fmt.Printf("%+v\n", err)
			} else if counter = counter + 1; counter < 5 {
				fmt.Printf("%+v\n", err)
			}
		}
		fmt.Printf("admins are %v\n", st.Admins())
		return fmt.Errorf("found %d errors", len(errs))
	} else if st != nil {
		fmt.Println(st)
	}

	log.Printf("[server] listening on %s\n", srv.Addr)
	return srv.ListenAndServe()
}

// Handler returns the adapter's handler.
func CorsHandler(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[cors] %s %q\n", r.Method, r.URL.Path)
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, HEAD, OPTIONS, POST, PUT")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			w.WriteHeader(http.StatusNoContent)
			return
		} else if r.Method == "GET" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, HEAD, OPTIONS, POST, PUT")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		}
		next.ServeHTTP(w, r)
	}
}
