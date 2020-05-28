package util

import (
	"math/rand"
	"time"
)

// 传入返回的字符串长度以及是否需要包含*这些特殊字符
func RandStr(n int, symbol bool) string {
	letter := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	if symbol {
		letter = append(letter, []rune("_*`~!@#$%^&()-+=|{}[]\\")...)
	}
	rand.Seed(time.Now().UnixNano())
	var s = make([]rune, n)
	for i := range s {
		s[i] = letter[rand.Intn(len(letter))]
	}
	return string(s)
}
