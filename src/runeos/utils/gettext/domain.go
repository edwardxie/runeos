// Copyright 2013 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gettext

import (
	"path/filepath"
	"strings"
)

type Domain struct {
	Languages map[string]*Catalog
}

// Create a new Domain by parsing .mo files from path which must have subdirectories
// matching the GLOB path/*/LC_MESSAGES/name.mo
func NewDomain(name string, path string) (*Domain, error) {
	files, err := filepath.Glob(filepath.Join(path, "*", "LC_MESSAGES", name+".mo"))
	if err != nil {
		return nil, err
	}
	domain := &Domain{
		Languages: make(map[string]*Catalog),
	}
	for _, f := range files {
		fs := strings.Split(f, "/")
		langCode := strings.ToLower(fs[len(fs)-3])
		domain.Languages[langCode] = &Catalog{}
	}
	return domain, nil
}

func (d *Domain) GetCatalog(langCode string) *Catalog {
	langCode = strings.ToLower(langCode)
	if c, _ := d.Languages[langCode]; c != nil {
		return c
	}
	// Fallback to just language if langCode includes country code
	i := strings.Index(langCode, "_")
	if i > 0 {
		langCode = langCode[:i]
		if c, _ := d.Languages[langCode]; c != nil {
			return c
		}
	}
	return nil
}

func (d *Domain) GetText(langCode string, text string) string {
	return d.GetCatalog(langCode).GetText(text)
}

func (d *Domain) NGetText(langCode string, singular string, plural string, n int) string {
	return d.GetCatalog(langCode).NGetText(singular, plural, n)
}
