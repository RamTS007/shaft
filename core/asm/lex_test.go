// Copyright 2017 The go-shaft Authors
// This file is part of the go-shaft library.
//
// The go-shaft library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-shaft library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-shaft library. If not, see <http://www.gnu.org/licenses/>.

package asm

import "testing"

func lexAll(src string) []token {
	ch := Lex("test.asm", []byte(src), false)

	var tokens []token
	for i := range ch {
		tokens = append(tokens, i)
	}
	return tokens
}

func TestComment(t *testing.T) {
	tokens := lexAll(";; this is a comment")
	if len(tokens) != 2 { // {new line, EOF}
		t.Error("expected no tokens")
	}
}
