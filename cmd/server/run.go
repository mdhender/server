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
	"github.com/mdhender/server/pkg/handlers/spa"
	"log"
	"net/http"
)

func run(cfg *config) error {
	var options []func(*server) error
	options = append(options, setSalt(cfg.Server.Salt))

	srv, err := newServer(cfg, options...)
	if err != nil {
		return err
	}
	srv.Handler = routes(srv, http.StripPrefix("/", spa.Handler(cfg.Server.PublicRoot)), cfg.Games.FileSavePath)

	log.Printf("[server] listening on %s\n", srv.Addr)
	return srv.ListenAndServe()
}
