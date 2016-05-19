package leftpad

import (
	"strings"
	"unicode/utf8"
)

// Leftpad ...
func Leftpad(s string, length int, ch ...rune) string {
	c := ' '
	if len(ch) > 0 {
		c = ch[0]
	}
	l := length - utf8.RuneCountInString(s)
	if l > 0 {
		s = strings.Repeat(string(c), l) + s
	}
	return s
}
