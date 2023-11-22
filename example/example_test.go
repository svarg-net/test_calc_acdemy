package example

import (
	"testing"
)

type testpair struct {
	values string
	result string
}

var tests = []testpair{
	{"1+1", "2"},
	{"9+1", "10"},
	{"I+III", "IV"},
	{"X+X", "XX"},
	{"IX+X", "XIX"},
}

func TestExample(t *testing.T) {
	for _, pair := range tests {
		ex := Example{StrExample: pair.values}
		ex.Construct()
		result := ex.Result()
		if result != pair.result {
			t.Error(
				"For", pair.values,
				"expected", pair.result,
				"got", result,
			)
		}
	}
}
