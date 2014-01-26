// Copyright 2013 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package gettext_test

import (
	"fmt"

	"code.google.com/p/gettext-go/gettext"
)

func ExampleLevenshteinString() {
	a := []rune("kitten")
	b := []rune("sitting")
	d := gettext.LevenshteinString(a, b)
	fmt.Println(d)
	// Output: 3
}

func ExampleLevenshteinWords() {
	a := []string{"What", "day", "is", "today", "?"}
	b := []string{"What", "is", "the", "date", "today", "?"}
	d := gettext.LevenshteinWords(a, b)
	fmt.Println(d)
	// Output: 3
}
