// Copyright 2013 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package gettext

import (
	"testing"
)

func TestLevenshteinString(t *testing.T) {
	if d := LevenshteinString([]rune("kitten"), []rune("sitting")); d != 3 {
		t.Fatalf("expect = %v, got = %v", 3, d)
	}
}

func TestLevenshteinWords(t *testing.T) {
	a := []string{"What", "day", "is", "today", "?"}
	b := []string{"What", "is", "the", "date", "today", "?"}
	if d := LevenshteinWords(a, b); d != 3 {
		t.Fatalf("expect = %v, got = %v", 3, d)
	}
}
