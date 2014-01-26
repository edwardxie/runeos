// Copyright 2013 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gettext

// GetLocale gets the program's current locale.
func GetLocale() string {
	return ""
}

// SetLocale sets the program's current locale.
func SetLocale(locale string) error {
	return nil
}

// GetTextdomain gets the current message domain.
func GetTextdomain(domainname string) string {
	return ""
}

// BindTextdomain sets directory containing message catalogs.
func BindTextdomain(domainname string, dirname string) string {
	return ""
}

// Attempt to translate a text string into the user's native language, by
// looking up the translation in a message catalog.
func Gettext(msgid string) string {
	return msgid
}

// Like Gettext(), but looking up the message in the specified domain.
func DGettext(domain string, msgid string) string {
	return msgid
}

// Attempt to translate a text string into the user's native language, by
// looking up the appropriate plural form of the translation in a message
// catalog.
func NGettext(msgid string, msgid_plural string, n uint64) string {
	return msgid
}

// Like NGettext(), but looking up the message in the specified domain.
func DNGettext(domainname string, msgid string, msgid_plural string, n uint64) string {
	return msgid
}
