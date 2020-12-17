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

// Package auth implements good-enough-for-free authorization bits.
package auth

type Authorization struct {
	ID    string          // id of the entity authorized
	Roles map[string]bool // map of the roles the entity has been authorized for
}

// HasAllRoles returns true if the entity has all of the given roles assigned to it.
func (a *Authorization) HasAllRole(roles ...string) bool {
	if a != nil && a.Roles != nil && len(roles) != 0 {
		for _, role := range roles {
			if !a.Roles[role] {
				return false
			}
		}
		return true
	}
	return false
}

// HasAnyRole returns true if the entity has any of the given roles assigned to it.
func (a *Authorization) HasAnyRole(roles ...string) bool {
	if a != nil && a.Roles != nil {
		for _, role := range roles {
			if a.Roles[role] {
				return true
			}
		}
	}
	return false
}

// HasRole returns true if the entity has the given roles assigned to it.
func (a *Authorization) HasRole(role string) bool {
	if a != nil && a.Roles != nil {
		return a.Roles[role]
	}
	return false
}
