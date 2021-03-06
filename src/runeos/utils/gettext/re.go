// Copyright 2013 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gettext

import (
	"regexp"
)

var (
	reComment                = regexp.MustCompile(`^\s*#`)            // #
	reExtractedComments      = regexp.MustCompile(`^\s*#\.`)          // #.
	reReferenceComments      = regexp.MustCompile(`^\s*#:`)           // #:
	reFlagsComments          = regexp.MustCompile(`^\s*#,`)           // #, fuzzy,c-format
	rePrevMsgContextComments = regexp.MustCompile(`^\s*#|\s+msgctxt`) // #| msgctxt
	rePrevMsgIdComments      = regexp.MustCompile(`^\s*#|\s+msgid`)   // #| msgid

	reMsgContext   = regexp.MustCompile(`^msgctxt\s+".*"\s*$`)            // msgctxt
	reMsgId        = regexp.MustCompile(`^msgid\s+".*"\s*$`)              // msgid
	reMsgIdPlural  = regexp.MustCompile(`^msgid_plural\s+".*"\s*$`)       // msgid_plural
	reMsgStr       = regexp.MustCompile(`^msgstr\s*".*"\s*$`)             // msgstr
	reMsgStrPlural = regexp.MustCompile(`^msgstr\s*(\[\d+\])\s*".*"\s*$`) // msgstr[0]
	reStringLine   = regexp.MustCompile(`^\s*".*"\s*$`)                   // "message"
	reBlankLine    = regexp.MustCompile(`^\s*$`)                          //
)

func (p *Message) isInvalidLine(s string) bool {
	if reComment.MatchString(s) {
		return false
	}
	if reBlankLine.MatchString(s) {
		return false
	}

	if reMsgContext.MatchString(s) {
		return false
	}
	if reMsgId.MatchString(s) {
		return false
	}
	if reMsgIdPlural.MatchString(s) {
		return false
	}
	if reMsgStr.MatchString(s) {
		return false
	}
	if reMsgStrPlural.MatchString(s) {
		return false
	}

	if reStringLine.MatchString(s) {
		return false
	}

	return true
}
