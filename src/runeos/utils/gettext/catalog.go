// Copyright 2013 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gettext

import (
	"net/textproto"
)

type Catalog struct {
	Header textproto.MIMEHeader
	Entrys map[string]Message
}

func (c *Catalog) GetText(text string) string {
	return text
}

func (c *Catalog) NGetText(singular string, plural string, n int) string {
	return singular
}
