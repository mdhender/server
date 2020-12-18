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

package memory

import "errors"

// ErrDuplicateAddress is used when the e-mail address is not unique.
var ErrDuplicateAddress = errors.New("duplicate e-mail address")

// ErrDuplicateID is used when the id is not unique.
var ErrDuplicateID = errors.New("duplicate id")

// ErrDuplicateName is used when the user name is not unique.
var ErrDuplicateName = errors.New("duplicate user name")

// ErrInvalidEmail is used when the email is not valid.
var ErrInvalidEmail = errors.New("invalid e-mail")

// ErrInvalidID is used when the id is not valid.
var ErrInvalidID = errors.New("invalid id")

// ErrInvalidName is used when the name is not valid.
var ErrInvalidName = errors.New("invalid name")
