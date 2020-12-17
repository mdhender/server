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
	"github.com/mdhender/server/pkg/http/rest"
	"github.com/mdhender/server/pkg/listing"
	"github.com/mdhender/server/pkg/way"
	"net/http"
)

func routes(s *server, rc routeConfig) http.Handler {
	router := way.NewRouter()

	router.Handle("GET", "/api/games", s.getAllGames())
	router.Handle("GET", "/api/games/:id", s.getGame())
	router.Handle("GET", "/api/games/:id/players", s.getGamePlayers())
	router.Handle("GET", "/api/games/:id/players/:playerName", s.getGamePlayer())
	router.Handle("GET", "/api/games/:id/players/:playerName/printout", s.getGamePlayerPrintout())
	router.Handle("GET", "/api/games/:id/systems", s.getGameSystems())
	router.Handle("GET", "/api/games/:id/systems/:systemId", s.getGameSystem())
	router.Handle("GET", "/api/users", rest.GetAllUsers(rc.services.userListing))
	router.Handle("GET", "/api/users/:id", rest.GetUser(rc.services.userListing))
	router.Handle("GET", "/api/version", s.getVersion())

	router.Handle("POST", "/api/engine/restart", s.restart())
	router.Handle("POST", "/api/games", s.addGame())
	router.Handle("POST", "/api/games/:id/orders", s.postGameOrders())
	router.Handle("POST", "/api/games/:id/save", s.postGameSave(rc.gameFileSavePath))
	router.Handle("POST", "/api/users", s.addUser())

	// assume that all other routes are to serve the front end application
	router.NotFound = rc.notFound

	return router
}

type routeConfig struct {
	gameFileSavePath string
	notFound         http.Handler
	services         struct {
		userListing listing.Service
	}
}
