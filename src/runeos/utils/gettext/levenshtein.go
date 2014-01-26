// Copyright 2013 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package gettext

// LevenshteinString calculate Levenshtein distance between two strings.
func LevenshteinString(s, t []rune) int {
	switch {
	case len(s) == 0:
		return len(t)
	case len(t) == 0:
		return len(s)
	case s[0] == t[0]:
		return LevenshteinString(s[1:], t[1:])
	}
	a := LevenshteinString(s[1:], t[1:])
	b := LevenshteinString(s, t[1:])
	c := LevenshteinString(s[1:], t)
	if a > b {
		a = b
	}
	if a > c {
		a = c
	}
	return a + 1
}

// LevenshteinWords calculate Levenshtein distance between two word list.
func LevenshteinWords(s, t []string) int {
	switch {
	case len(s) == 0:
		return len(t)
	case len(t) == 0:
		return len(s)
	case s[0] == t[0]:
		return LevenshteinWords(s[1:], t[1:])
	}
	a := LevenshteinWords(s[1:], t[1:])
	b := LevenshteinWords(s, t[1:])
	c := LevenshteinWords(s[1:], t)
	if a > b {
		a = b
	}
	if a > c {
		a = c
	}
	return a + 1
}
