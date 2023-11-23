package example

import (
	"strconv"
)

const (
	TypeNumbers1 = "Arabic"
	TypeNumbers2 = "Rome"
)

type Number struct {
	Value      string
	Int        int
	TypeNumber string
}

func (n *Number) getInt() {
	value, _ := strconv.Atoi(n.Value)
	if n.Value != "0" && value == 0 {
		n.TypeNumber = TypeNumbers2
		n.romeToArabic()
	} else {
		n.TypeNumber = TypeNumbers1
		n.Int = value
	}
}

type RomanNumeral struct {
	Value  int
	Symbol string
}

var RomanNumerals = []RomanNumeral{
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func (n *Number) romeToArabic() {
	result := 0
	for i, v := range n.Value {
		val := findSymbol(RomanNumerals, string(v))
		valNext := 0
		if i+1 < len(n.Value) {
			valNext = findSymbol(RomanNumerals, string(n.Value[i+1]))
		}
		if i < len(n.Value)-1 && val < valNext {
			result -= val
		} else {
			result += val
		}
	}
	n.Int = result
}

func (n *Number) toRome() {
	result := ""
	for _, r := range RomanNumerals {
		for n.Int >= r.Value {
			result += r.Symbol
			n.Int -= r.Value
		}
	}
	n.Value = result
}

func findSymbol(romanNumerals []RomanNumeral, symbol string) int {
	for _, r := range romanNumerals {
		if r.Symbol == symbol {
			return r.Value
		}
	}
	return 0
}
