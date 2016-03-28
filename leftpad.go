package leftpad

import "strings"

// Leftpad ...
func Leftpad(s string, length int, ch ...rune) string {
	c := ' '
	if len(ch) > 0 {
		c = ch[0]
	}
	l := length - len(s)
	if l > 0 {
		s = strings.Repeat(string(c), l) + s
	}
	return s
}
