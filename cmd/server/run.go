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
