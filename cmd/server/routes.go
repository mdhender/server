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
	"github.com/mdhender/server/internal/obsolete/adding"
	"github.com/mdhender/server/internal/obsolete/http/rest"
	"github.com/mdhender/server/internal/obsolete/listing"
	"github.com/mdhender/server/internal/obsolete/reporting"
	"github.com/mdhender/server/internal/obsolete/updating"
	"github.com/mdhender/server/internal/way"
	"net/http"
)

func routes(s *server, rc routeConfig) http.Handler {
	router := way.NewRouter()

	router.Handle("GET", "/api/game/:id", rest.GetGame(rc.services.listing))
	router.Handle("GET", "/api/game/:id/player/:player_name", rest.GetGamePlayer(rc.services.listing))
	router.Handle("GET", "/api/game/:id/player/:player_name/print-out", rest.GetGamePlayerPrintout(rc.services.reporting))
	router.Handle("GET", "/api/game/:id/player/:player_name/print-out/turn/:turn_number", rest.GetGamePlayerPrintout(rc.services.reporting))
	router.Handle("GET", "/api/game/:id/players", rest.GetGamePlayers(rc.services.listing))
	router.Handle("GET", "/api/game/:id/system/:system_name", rest.GetGameSystem(rc.services.listing))
	router.Handle("GET", "/api/game/:id/systems", rest.GetGameSystems(rc.services.listing))
	router.Handle("GET", "/api/games", rest.GetGames(rc.services.listing))
	router.Handle("GET", "/api/user/:id", rest.GetUser(rc.services.listing))
	router.Handle("GET", "/api/users", rest.GetUsers(rc.services.listing))
	router.Handle("GET", "/api/version", rest.GetVersion(rc.services.listing))
	router.Handle("GET", "/api/frak", frak())

	router.Handle("POST", "/api/engine/restart", s.restart())
	router.Handle("POST", "/api/game/orders", rest.UpdateGameOrders(rc.services.updating))
	router.Handle("POST", "/api/game/save", rest.UpdateGame(rc.services.updating))
	router.Handle("POST", "/api/games/create", rest.AddGame(rc.services.adding))
	router.Handle("POST", "/api/users/create", rest.AddUser(rc.services.adding))

	return router
}

type routeConfig struct {
	gameFileSavePath string
	services         struct {
		adding    adding.Service
		listing   listing.Service
		reporting reporting.Service
		updating  updating.Service
	}
}

func frak() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
