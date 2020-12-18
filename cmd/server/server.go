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
	"crypto/md5"
	"encoding/binary"
	"github.com/mdhender/server"
	"github.com/mdhender/server/internal/gamemeta"
	"github.com/mdhender/server/internal/users"
	"io"
	"net"
	"net/http"
)

var versionMajor string = "0"
var versionMinor string = "0"
var versionPatch string = "0"

// server defines the server
type server struct {
	http.Server
	salt  string
	turns []engine.Orders
	games map[string]*gamemeta.GameMeta
	users *users.Users
}

// serverContextKey is the context key type for storing parameters in context.Context.
type serverContextKey string

// newServer returns an initialized server.
// the main change from the default server is that we override the default timeouts.
// see the following sources for an explanation of why:
//   https://blog.cloudflare.com/exposing-go-on-the-internet/
//   https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
//   https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
func newServer(cfg *config, options ...func(*server) error) (s *server, err error) {
	s = &server{}
	s.Addr = net.JoinHostPort(cfg.Server.Host, cfg.Server.Port)
	s.IdleTimeout = cfg.Server.Timeout.Idle
	s.ReadTimeout = cfg.Server.Timeout.Read
	s.WriteTimeout = cfg.Server.Timeout.Write
	s.MaxHeaderBytes = 1 << 20

	// allow caller to override the default values
	for _, option := range options {
		if err := option(s); err != nil {
			return nil, err
		}
	}

	return s, nil
}

func setSalt(salt string) func(*server) error {
	return func(s *server) error {
		s.salt = salt
		return nil
	}
}

func (s *server) seed(seedString string) int64 {
	switch seedString {
	case "1812":
		return 1812
	case "1917":
		return 1917
	}
	hasher := md5.New()
	io.WriteString(hasher, s.salt)
	io.WriteString(hasher, seedString)
	return int64(binary.BigEndian.Uint64(hasher.Sum(nil)))
}
