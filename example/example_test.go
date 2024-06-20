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
	{"9*9", "81"},
	{"10*10", "100"},
	{"10/9", "1"},
	{"1-10", "-9"},
	{"I+III", "IV"},
	{"X+X", "XX"},
	{"IX+X", "XIX"},
	//{"X-X", "Римское число не должно быть отрицательным"},
	//{"IX-X", "Римское число не должно быть отрицательным"},
	//{"11-10", "Калькулятор принимает на вход числа от 1 до 10"},
	//{"1-12", "Калькулятор принимает на вход числа от 1 до 10"},
	//{"14-13", "Калькулятор принимает на вход числа от 1 до 10"},
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
