package utils

import "testing"

func TestReverse(t *testing.T) {
	cases := [] struct{
		input, expected string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		actual := Reverse(c.input)
		if actual != c.expected {
			t.Errorf("Reverse(%q) == %q, expectd %q", c.input, actual, c.expected)
		}
	}

}
