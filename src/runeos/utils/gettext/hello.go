// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"
	"log"

	"code.google.com/p/gettext-go/gettext"
)

func main() {
	poFile, err := gettext.LoadPoFile("../testdata/test.po")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", poFile)
}
