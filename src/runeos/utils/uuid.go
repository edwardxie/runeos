/*
一组UUID，是由一串16位组（亦称128位）的16进位数字所构成，是故UUID理论上的总数为
216 x 8=2128，约等于3.4 x 1038。也就是说若每纳秒产生1兆个UUID，要花100亿年才
会将所有UUID用完。
UUID的标准型式包含32个16进位数字，以连字号分为五段，形式为8-4-4-4-12的36个字符。
示例：550e8400-e29b-41d4-a716-446655440000

UUID亦可刻意重复以表示同类。例如说微软的COM中，所有组件皆必须实现出IUnknown接口，
方法是产生一个代表IUnknown的UUID。无论是程序试图访问组件中的IUnknown接口，或是
实现IUnknown接口的组件，只要IUnknown一被使用，皆会被参考至同一个ID：
00000000-0000-0000-C000-000000000046。
*/

package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

const (
	defaultDigit = 16
)

type uuid struct {
	digit  int
	random []byte
}

type Identifier interface {
	MD5() []byte
	URL(int) []byte
	UUID() string
}

func NewID() Identifier {
	r := &uuid{}
	return r
}

func (u *uuid) ini() {
	u.random = make([]byte, u.digit)
	io.ReadFull(rand.Reader, u.random)
}

func (u *uuid) MD5() []byte {
	u.digit = defaultDigit
	u.ini()
	return md5.New().Sum(u.random)
}

func (u *uuid) URL(digit int) []byte {
	u.digit = digit
	u.ini()
	return []byte(base64.URLEncoding.EncodeToString(u.random))
}

func (u *uuid) UUID() string {
	u.digit = defaultDigit
	u.ini()
	r := md5.New().Sum(u.random)
	return fmt.Sprintf("%x-%x-%x-%x-%x", r[:4], r[4:6], r[6:8], r[8:10], r[10:16])
}
