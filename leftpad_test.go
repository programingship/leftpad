package leftpad

import "testing"

// TestLeftpad ...
func TestLeftpad(t *testing.T) {
	ts := []struct {
		s, expected string
	}{
		{Leftpad("foo", 5), "  foo"},
		{Leftpad("foobar", 6), "foobar"},
		{Leftpad("1", 2, '0'), "01"},
		{Leftpad("好吗", 3, '你'), "你好吗"},
	}
	for _, v := range ts {
		if v.expected != v.s {
			t.Errorf("Expected %s, actual %s", v.expected, v.s)
		}
	}
}
