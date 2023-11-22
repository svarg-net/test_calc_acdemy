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

func (n *Number) romeToArabic() {
	romeNumerals := map[string]int{"X": 10, "V": 5, "I": 1}
	result := 0
	for i, v := range n.Value {
		if i < len(n.Value)-1 && romeNumerals[string(v)] < romeNumerals[string(n.Value[i+1])] {
			result -= romeNumerals[string(v)]
		} else {
			result += romeNumerals[string(v)]
		}
	}
	n.Int = result

}

func (n *Number) toRome() {
	romeNumerals := map[int]string{10: "X", 9: "IX", 5: "V", 4: "IV", 1: "I"}
	result := ""
	for n.Int > 0 {
		for i, r := range romeNumerals {
			if n.Int > i {
				result += r
				n.Int -= i
			} else {

			}
		}
	}
	n.Value = result

}
