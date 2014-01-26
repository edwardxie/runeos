// Copyright 2013 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gettext

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Comment struct {
	TranslatorComment string   // #  translator-comments // TrimSpace
	ExtractedComment  string   // #. extracted-comments
	ReferenceFile     []string // #: src/msgcmp.c:338 src/po-lex.c:699
	ReferenceLine     []int    // #: src/msgcmp.c:338 src/po-lex.c:699
	Flags             []string // #, fuzzy,c-format,range:0..10
	PrevMsgContext    string   // #| msgctxt previous-context
	PrevMsgId         string   // #| msgid previous-untranslated-string
}

func (p *Comment) readPoComment(r *lineReader) (err error) {
	*p = Comment{}
	if err = r.skipBlankLine(); err != nil {
		return err
	}
	defer func(oldPos int) {
		newPos := r.currentPos()
		if newPos != oldPos && err == io.EOF {
			err = nil
		}
	}(r.currentPos())

	for {
		var s string
		if s, _, err = r.currentLine(); err != nil {
			return
		}
		if len(s) == 0 || s[0] != '#' {
			return
		}

		if err = p.readTranslatorComment(r); err != nil {
			return
		}
		if err = p.readExtractedComment(r); err != nil {
			return
		}
		if err = p.readReferenceComment(r); err != nil {
			return
		}
		if err = p.readFlagsComment(r); err != nil {
			return
		}
		if err = p.readPrevMsgContext(r); err != nil {
			return
		}
		if err = p.readPrevMsgId(r); err != nil {
			return
		}
	}

	return nil
}

func (p *Comment) readTranslatorComment(r *lineReader) (err error) {
	const prefix = "# " // .,:|
	for {
		var s string
		if s, _, err = r.readLine(); err != nil {
			return err
		}
		if len(s) < 1 || s[0] != '#' {
			r.unreadLine()
			return nil
		}
		if len(s) >= 2 {
			switch s[1] {
			case '.', ',', ':', '|':
				r.unreadLine()
				return nil
			}
		}
		if p.TranslatorComment != "" {
			p.TranslatorComment += "\n"
		}
		p.TranslatorComment += strings.TrimSpace(s[1:])
	}
	return nil
}

func (p *Comment) readExtractedComment(r *lineReader) (err error) {
	const prefix = "#. "
	for {
		var s string
		if s, _, err = r.readLine(); err != nil {
			return err
		}
		if len(s) < len(prefix) || s[:len(prefix)] != prefix {
			r.unreadLine()
			return nil
		}
		if p.ExtractedComment != "" {
			p.ExtractedComment += "\n"
		}
		p.ExtractedComment += s[len(prefix):]
	}
	return nil
}

func (p *Comment) readReferenceComment(r *lineReader) (err error) {
	const prefix = "#: "
	for {
		var s string
		if s, _, err = r.readLine(); err != nil {
			return err
		}
		if len(s) < len(prefix) || s[:len(prefix)] != prefix {
			r.unreadLine()
			return nil
		}
		ss := strings.Split(strings.TrimSpace(s[len(prefix):]), " ")
		for i := 0; i < len(ss); i++ {
			idx := strings.Index(ss[i], ":")
			if idx <= 0 {
				continue
			}
			name := strings.TrimSpace(ss[i][:idx])
			line, _ := strconv.Atoi(strings.TrimSpace(ss[i][idx+1:]))
			p.ReferenceFile = append(p.ReferenceFile, name)
			p.ReferenceLine = append(p.ReferenceLine, line)
		}
	}
	return nil
}

func (p *Comment) readFlagsComment(r *lineReader) (err error) {
	const prefix = "#, "
	for {
		var s string
		if s, _, err = r.readLine(); err != nil {
			return err
		}
		if len(s) < len(prefix) || s[:len(prefix)] != prefix {
			r.unreadLine()
			return nil
		}
		ss := strings.Split(strings.TrimSpace(s[len(prefix):]), ",")
		for i := 0; i < len(ss); i++ {
			p.Flags = append(p.Flags, strings.TrimSpace(ss[i]))
		}
	}
	return nil
}

func (p *Comment) readPrevMsgContext(r *lineReader) (err error) {
	const prefix = "#| msgctxt "
	for {
		var s string
		if s, _, err = r.readLine(); err != nil {
			return err
		}
		if len(s) < len(prefix) || s[:len(prefix)] != prefix {
			r.unreadLine()
			return nil
		}
		if p.PrevMsgContext != "" {
			p.PrevMsgContext += "\n"
		}
		p.PrevMsgContext += s[len(prefix):]
	}
	return nil
}

func (p *Comment) readPrevMsgId(r *lineReader) (err error) {
	const prefix = "#| msgid "
	for {
		var s string
		if s, _, err = r.readLine(); err != nil {
			return err
		}
		if len(s) < len(prefix) || s[:len(prefix)] != prefix {
			r.unreadLine()
			return nil
		}
		if p.PrevMsgId != "" {
			p.PrevMsgId += "\n"
		}
		p.PrevMsgId += s[len(prefix):]
	}
	return nil
}

func (p *Comment) IsFuzzy() bool {
	for _, s := range p.Flags {
		if s == "fuzzy" {
			return true
		}
	}
	return false
}

func (p Comment) String() string {
	var buf bytes.Buffer
	if p.TranslatorComment != "" {
		ss := strings.Split(p.TranslatorComment, "\n")
		for i := 0; i < len(ss); i++ {
			fmt.Fprintf(&buf, "# %s\n", ss[i])
		}
	}
	if p.ExtractedComment != "" {
		ss := strings.Split(p.ExtractedComment, "\n")
		for i := 0; i < len(ss); i++ {
			fmt.Fprintf(&buf, "#. %s\n", ss[i])
		}
	}
	if a, b := len(p.ReferenceFile), len(p.ReferenceLine); a != 0 && a == b {
		fmt.Fprintf(&buf, "#:")
		for i := 0; i < len(p.ReferenceFile); i++ {
			fmt.Fprintf(&buf, " %s:%d", p.ReferenceFile[i], p.ReferenceLine[i])
		}
		fmt.Fprintf(&buf, "\n")
	}
	if len(p.Flags) != 0 {
		fmt.Fprintf(&buf, "#, %s", p.Flags[0])
		for i := 1; i < len(p.Flags); i++ {
			fmt.Fprintf(&buf, ", %s", p.Flags[i])
		}
		fmt.Fprintf(&buf, "\n")
	}
	if p.PrevMsgContext != "" {
		ss := strings.Split(p.PrevMsgContext, "\n")
		for i := 0; i < len(ss); i++ {
			fmt.Fprintf(&buf, "#| msgctxt %s\n", ss[i])
		}
	}
	if p.PrevMsgId != "" {
		ss := strings.Split(p.PrevMsgId, "\n")
		for i := 0; i < len(ss); i++ {
			fmt.Fprintf(&buf, "#| msgid %s\n", ss[i])
		}
	}
	return buf.String()
}
