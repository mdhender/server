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
	"fmt"
	engine "github.com/mdhender/server"
	"github.com/mdhender/server/pkg/jsonapi"
	"github.com/mdhender/server/pkg/way"
	"log"
	"net/http"
)

func routes(s *server) http.Handler {
	router := way.NewRouter()

	router.Handle("POST", "/api/orders", s.postOrders())
	router.Handle("POST", "/api/turn", s.postTurn())

	return router
}

// postOrders .
func (s *server) postOrders() http.HandlerFunc {
	type response struct {
		Message string `json:"msg"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		// Enforce a maximum read of 1MB from the request body.
		orders, errs := engine.Decode(http.MaxBytesReader(w, r.Body, 1<<20))
		if errs != nil {
			for _, err := range errs {
				log.Printf("[orders] %+v\n", err)
			}
			jsonapi.Error(w, r, http.StatusBadRequest, errs[0])
			return
		}
		log.Printf("[orders] %d %v\n", orders.Len(), orders)

		// do something with the orders
		var e engine.Engine
		if err := e.PostOrders(orders); err != nil {
			jsonapi.Error(w, r, http.StatusInternalServerError, err)
			return
		}

		jsonapi.Ok(w, r, http.StatusOK, response{"order accepted"})
	}
}

// postTurn .
func (s *server) postTurn() http.HandlerFunc {
	type response struct {
		TurnNo  int      `json:"turn"`
		Message string   `json:"msg"`
		Errors  []string `json:"errors,omitempty"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if len(s.turns) == 0 {
			jsonapi.Error(w, r, http.StatusBadRequest, fmt.Errorf("no turns to process"))
			return
		}
		if err := s.process(len(s.turns), s.turns[len(s.turns)-1]); err != nil {
			jsonapi.Error(w, r, http.StatusInternalServerError, err)
			return
		}
		jsonapi.Ok(w, r, http.StatusOK, response{TurnNo: len(s.turns), Message: "turn completed"})
	}
}

func (s *server) process(turnNo int, orders engine.Orders) error {
	return fmt.Errorf("server.process not implemented")
}
