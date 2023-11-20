package example

import (
	"log"
	"regexp"
	"strconv"
)

type ExampleI interface {
	Construct(string) ExampleStruct
}

type Example struct {
	StrExample string
	data       ExampleStruct
}

type ExampleStruct struct {
	FirstNumber  Number
	SecondNumber Number
	Action       string
}

func (e *Example) Construct() {
	e.getNumber()
	e.getAction()
	e.validation()
}

func (e *Example) validation() {
	if e.data.FirstNumber.TypeNumber != e.data.SecondNumber.TypeNumber {
		log.Fatal("Вывод ошибки, так как используются одновременно разные системы счисления")
	}
	if (e.data.FirstNumber.Int > 10 && e.data.FirstNumber.Int < 1) || (e.data.SecondNumber.Int > 10 && e.data.SecondNumber.Int < 1) {
		log.Fatal("Калькулятор принимает на вход числа от 1 до 10")
	}
}

func (e *Example) Result() string {
	FirstNumber := e.data.FirstNumber.Int
	SecondNumber := e.data.SecondNumber.Int
	resultNumber := Number{}
	switch e.data.Action {
	case "+":
		resultNumber.Int = (FirstNumber + SecondNumber)
	case "-":
		resultNumber.Int = (FirstNumber - SecondNumber)
	case "*":
		resultNumber.Int = (FirstNumber * SecondNumber)
	case "/":
		resultNumber.Int = (FirstNumber / SecondNumber)
	}
	if e.data.FirstNumber.TypeNumber == TypeNumbers2 && e.data.FirstNumber.TypeNumber == TypeNumbers2 {
		if resultNumber.Int <= 0 {
			return "Римское число не должно быть отрицательным"
		} else {
			resultNumber.toRome()
			return resultNumber.Value
		}
	} else {
		return strconv.Itoa(resultNumber.Int)
	}

}

func (e *Example) getAction() {
	re := regexp.MustCompile(`([0-9IVX]+)([\+\-\*/]{1})([0-9IVX]+)`)
	exampleNew := re.FindStringSubmatch(e.StrExample)
	e.data.Action = exampleNew[2]
}

func (e *Example) getNumber() {
	re := regexp.MustCompile(`[\+\-\*/]`)
	exampleNew := re.Split(e.StrExample, -1)
	if len(exampleNew) < 2 {
		log.Fatal("Вывод ошибки, так как строка не является математической операцией.")
	}

	if len(exampleNew) > 2 {
		log.Fatal("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}
	e.data.FirstNumber.Value = exampleNew[0]
	e.data.FirstNumber.getInt()
	e.data.SecondNumber.Value = exampleNew[1]
	e.data.SecondNumber.getInt()

}
