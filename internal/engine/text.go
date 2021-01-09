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

package engine

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Text tries to safely handle tainted input.
// By "tainted" we mean untrusted user input.
type Text struct {
	untainted bool
	text      string
}

func NewText(text string) (Text, error) {
	if !utf8.ValidString(text) {
		return Text{}, fmt.Errorf("invalid utf-8: %w", ERRBADREQUEST)
	}
	return Text{
		text: text,
	}, nil
}

func (t Text) Length() int {
	return len(t.text)
}

func (t Text) String() string {
	if !t.untainted {
		return "***tainted***"
	}
	return t.text
}

func (t Text) Text() string {
	return t.text
}

func (t Text) TrimSpace() Text {
	return Text{
		untainted: t.untainted,
		text:      strings.TrimSpace(t.text),
	}
}

func (t Text) Taint() Text {
	return Text{
		untainted: false,
		text:      t.text,
	}
}

func (t Text) Tainted() bool {
	return !t.untainted
}

func (t Text) Untaint() Text {
	return Text{
		untainted: true,
		text:      t.text,
	}
}

// sanitize is an attempt to replace problematic characters with an underscore.
// it also forces the string to be valid utf-8.'
// for some reason, it also avoids runs of replacement characters.
func sanitize(s string) string {
	var dst, prior string
	for src := []byte(s); len(src) != 0; {
		r, w := utf8.DecodeRune(src)
		switch r {
		case utf8.RuneError:
			if prior != " " {
				dst, prior = dst+" ", " "
			}
		case '\\', '<', '>', '%':
			if prior != "_" {
				dst, prior = dst+"_", "_"
			}
		default:
			if unicode.IsPrint(r) {
				dst += string(r)
			} else if unicode.IsSpace(r) {
				if prior != " " {
					dst, prior = dst+" ", " "
				}
			} else if prior != "_" {
				dst, prior = dst+"_", "_"
			}
		}
		src = src[w:]
	}
	return dst
}
