package utils

import "testing"

func TestSplit(t *testing.T) {
	cases := []struct{
		input int
		expected []int
	}{
		{10, {4, 6}},
		{-10, {-4, -6}},
		{0, {0, 0}},
	}
	for _, c := range cases {
		actual1, actual2 := Split(c.input)
		if actual1 != c.expected[0] || actual2 != c.expected[1] {
			t.Errorf("Split(%q) == %q, %q, expectd %q", c.input, actual1, actual2, c.expected)
		}
	}
}
