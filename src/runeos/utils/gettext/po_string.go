// Copyright 2013 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gettext

import (
	"bytes"
	"strings"
)

// DecodePoString decodes text from PO string format.
func DecodePoString(text string) string {
	lines := strings.Split(text, "\n")
	for i := 0; i < len(lines); i++ {
		left := strings.Index(lines[i], `"`)
		right := strings.LastIndex(lines[i], `"`)
		if left < 0 || right < 0 || left == right {
			lines[i] = ""
			continue
		}
		line := lines[i][left+1 : right]
		line = strings.Replace(line, `\"`, `"`, -1)
		line = strings.Replace(line, `\n`, "\n", -1)
		line = strings.Replace(line, `\\`, `\`, -1)
		lines[i] = line
	}
	return strings.Join(lines, "")
}

// EncodePoString encodes text to PO string format.
func EncodePoString(text string) string {
	var buf bytes.Buffer
	lines := strings.Split(text, "\n")
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			if i != len(lines)-1 {
				buf.WriteString(`"\n"` + "\n")
			}
			continue
		}
		buf.WriteRune('"')
		for _, r := range lines[i] {
			switch r {
			case '\\':
				buf.WriteString(`\\`)
			case '"':
				buf.WriteString(`\"`)
			case '\n':
				buf.WriteString(`\n`)
			default:
				buf.WriteRune(r)
			}
		}
		buf.WriteString(`\n"` + "\n")
	}
	return buf.String()
}
