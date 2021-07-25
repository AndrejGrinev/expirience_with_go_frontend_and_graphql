package utils

import "testing"

type testpair struct {
	value int
	res   string
}

var tests = []testpair{
	{1, "1"},
	{-2, "-2"},
	{99, "99"},
}

func TestToStr(t *testing.T) {
	for _, pair := range tests {
		v := ToStr(pair.value)
		if v != pair.res {
			t.Error(
				"For", pair.value,
				"expected", pair.res,
				"got", v,
			)
		}
	}
}
