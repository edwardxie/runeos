// Copyright 2013 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gettext

import (
	"reflect"
	"testing"
)

var (
	testPoEditPoFile = "../testdata/poedit-1.5.7-zh_CN.po"
	testPoEditMoFile = "../testdata/poedit-1.5.7-zh_CN.mo"
)

func _TestPoEditPoFile(t *testing.T) {
	po, err := LoadPoFile(testPoEditPoFile)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(&po.MimeHeader, &poEditFile.MimeHeader) {
		t.Fatalf("expect = %v, got = %v", &poEditFile.MimeHeader, &po.MimeHeader)
	}
	for i := 0; i < len(po.Messages) && i < len(poEditFile.Messages); i++ {
		if !reflect.DeepEqual(&po.Messages[i], &poEditFile.Messages[i]) {
			t.Fatalf("%d: expect = %v, got = %v", i, poEditFile.Messages[i], po.Messages[i])
		}
	}
}

var poEditFile = &File{}
