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
	"github.com/mdhender/server/pkg/creating"
	"github.com/mdhender/server/pkg/http/rest"
	"github.com/mdhender/server/pkg/listing"
	"github.com/mdhender/server/pkg/way"
	"net/http"
)

func routes(s *server, rc routeConfig) http.Handler {
	router := way.NewRouter()

	router.Handle("GET", "/api/games", rest.GetGames(rc.services.listing))
	router.Handle("GET", "/api/games/:id", rest.GetGame(rc.services.listing))
	router.Handle("GET", "/api/games/:id/players", rest.GetGamePlayers(rc.services.listing))
	router.Handle("GET", "/api/games/:id/players/:player_name", rest.GetGamePlayer(rc.services.listing))
	router.Handle("GET", "/api/games/:id/players/:player_name/printout", s.getGamePlayerPrintout())
	router.Handle("GET", "/api/games/:id/systems", s.getGameSystems())
	router.Handle("GET", "/api/games/:id/systems/:systemId", s.getGameSystem())
	router.Handle("GET", "/api/users", rest.GetUsers(rc.services.listing))
	router.Handle("GET", "/api/users/:id", rest.GetUser(rc.services.listing))
	router.Handle("GET", "/api/version", s.getVersion())

	router.Handle("POST", "/api/engine/restart", s.restart())
	router.Handle("POST", "/api/games", s.addGame())
	router.Handle("POST", "/api/games/:id/orders", s.postGameOrders())
	router.Handle("POST", "/api/games/:id/save", s.postGameSave(rc.gameFileSavePath))
	router.Handle("POST", "/api/users", rest.CreateUser(rc.services.creating))

	// assume that all other routes are to serve the front end application
	router.NotFound = rc.notFound

	return router
}

type routeConfig struct {
	gameFileSavePath string
	notFound         http.Handler
	services         struct {
		creating creating.Service
		listing  listing.Service
	}
}
