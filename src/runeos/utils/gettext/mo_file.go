// Copyright 2013 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gettext

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	moMagicLittleEndian = 0x950412de
	moMagicBigEndian    = 0xde120495

	nulSeparator = "\x00"
	eotSeparator = "\x04"
)

type _MoFile struct {
	MajorVersion uint16
	MinorVersion uint16
	MimeHeader   map[string]string
	MsgStrTable  map[string]_MoEntry
}

type _MoEntry struct {
	MsgContext   string   // msgctxt context
	MsgId        string   // msgid untranslated-string
	MsgIdPlural  string   // msgid_plural untranslated-string-plural
	MsgStr       string   // msgstr translated-string
	MsgStrPlural []string // msgstr[0] translated-string-case-0
}

func _LoadMoFile(name string) (*_MoFile, error) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return _LoadMoData(data)
}

func _LoadMoData(data []byte) (*_MoFile, error) {
	r := bytes.NewReader(data)

	var magicNumber uint32
	if err := binary.Read(r, binary.LittleEndian, &magicNumber); err != nil {
		return nil, fmt.Errorf("gettext: %v", err)
	}
	var bo binary.ByteOrder
	switch magicNumber {
	case moMagicLittleEndian:
		bo = binary.LittleEndian
	case moMagicBigEndian:
		bo = binary.BigEndian
	default:
		return nil, fmt.Errorf("gettext: %v", "invalid magic number")
	}

	var header struct {
		MajorVersion uint16
		MinorVersion uint16
		MsgIdCount   uint32
		MsgIdOffset  uint32
		MsgStrOffset uint32
		HashSize     uint32
		HashOffset   uint32
	}
	if err := binary.Read(r, bo, &header); err != nil {
		return nil, fmt.Errorf("gettext: %v", err)
	}
	if v := header.MajorVersion; v != 0 && v != 1 {
		return nil, fmt.Errorf("gettext: %v", "invalid version number")
	}
	if v := header.MinorVersion; v != 0 && v != 1 {
		return nil, fmt.Errorf("gettext: %v", "invalid version number")
	}

	msgIdStart := make([]uint32, header.MsgIdCount)
	msgIdLen := make([]uint32, header.MsgIdCount)
	if _, err := r.Seek(int64(header.MsgIdOffset), 0); err != nil {
		return nil, fmt.Errorf("gettext: %v", err)
	}
	for i := 0; i < int(header.MsgIdCount); i++ {
		if err := binary.Read(r, bo, &msgIdLen[i]); err != nil {
			return nil, fmt.Errorf("gettext: %v", err)
		}
		if err := binary.Read(r, bo, &msgIdStart[i]); err != nil {
			return nil, fmt.Errorf("gettext: %v", err)
		}
	}

	msgStrStart := make([]int32, header.MsgIdCount)
	msgStrLen := make([]int32, header.MsgIdCount)
	if _, err := r.Seek(int64(header.MsgStrOffset), 0); err != nil {
		return nil, fmt.Errorf("gettext: %v", err)
	}
	for i := 0; i < int(header.MsgIdCount); i++ {
		if err := binary.Read(r, bo, &msgStrLen[i]); err != nil {
			return nil, fmt.Errorf("gettext: %v", err)
		}
		if err := binary.Read(r, bo, &msgStrStart[i]); err != nil {
			return nil, fmt.Errorf("gettext: %v", err)
		}
	}

	file := &_MoFile{
		MajorVersion: header.MajorVersion,
		MinorVersion: header.MinorVersion,
		MimeHeader:   make(map[string]string),
		MsgStrTable:  make(map[string]_MoEntry),
	}
	for i := 0; i < int(header.MsgIdCount); i++ {
		if _, err := r.Seek(int64(msgIdStart[i]), 0); err != nil {
			return nil, fmt.Errorf("gettext: %v", err)
		}
		msgIdData := make([]byte, msgIdLen[i])
		if _, err := r.Read(msgIdData); err != nil {
			return nil, fmt.Errorf("gettext: %v", err)
		}

		if _, err := r.Seek(int64(msgStrStart[i]), 0); err != nil {
			return nil, fmt.Errorf("gettext: %v", err)
		}
		msgStrData := make([]byte, msgStrLen[i])
		if _, err := r.Read(msgStrData); err != nil {
			return nil, fmt.Errorf("gettext: %v", err)
		}

		if len(msgIdData) == 0 {
			ss := strings.Split(string(msgStrData), "\n")
			for i := 0; i < len(ss); i++ {
				idx := strings.Index(ss[i], ":")
				if idx < 0 {
					continue
				}
				key := strings.TrimSpace(ss[i][:idx])
				val := strings.TrimSpace(ss[i][idx+1:])
				file.MimeHeader[key] = val
			}
		} else {
			var msg = _MoEntry{
				MsgId:  string(msgIdData),
				MsgStr: string(msgStrData),
			}
			// Is this a context message?
			if idx := strings.Index(msg.MsgId, eotSeparator); idx != -1 {
				msg.MsgContext, msg.MsgId = msg.MsgId[:idx], msg.MsgId[idx+1:]
			}
			// Is this a plural message?
			if idx := strings.Index(msg.MsgId, nulSeparator); idx != -1 {
				msg.MsgId, msg.MsgIdPlural = msg.MsgId[:idx], msg.MsgId[idx+1:]
				msg.MsgStrPlural = strings.Split(msg.MsgStr, nulSeparator)
				msg.MsgStr = ""
			}
			file.MsgStrTable[msg.MsgId] = msg
		}
	}

	return file, nil
}

func (f *_MoFile) Save(name string) error {
	return ioutil.WriteFile(name, f.Data(name), 0666)
}

// TODO
func (f *_MoFile) Data(name string) []byte {
	var buf bytes.Buffer

	var magicNumber = uint32(moMagicLittleEndian)
	binary.Write(&buf, binary.LittleEndian, &magicNumber)
	binary.Write(&buf, binary.LittleEndian, &f.MajorVersion)
	binary.Write(&buf, binary.LittleEndian, &f.MinorVersion)

	var strCount = uint32(len(f.MsgStrTable)) + 1
	_ = strCount

	return buf.Bytes()
}

func (f *_MoFile) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "# version: %d.%d\n", f.MajorVersion, f.MinorVersion)
	fmt.Fprintf(&buf, `msgid ""`+"\n")
	fmt.Fprintf(&buf, `msgstr ""`+"\n")
	for k, v := range f.MimeHeader {
		fmt.Fprintf(&buf, `"%s: %s\n"`+"\n", k, v)
	}
	fmt.Fprintf(&buf, "\n")

	for k, v := range f.MsgStrTable {
		fmt.Fprintf(&buf, `msgid "%s"`+"\n", k)
		fmt.Fprintf(&buf, `msgstr "%s"`+"\n", v.MsgStr)
		fmt.Fprintf(&buf, "\n")
	}

	return buf.String()
}
