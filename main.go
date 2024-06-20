package main

import (
	"academy/example"
	"bufio"
	"fmt"
	"os"
)

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Printf("%v", r)
	}
}

func main() {
	defer handlePanic()
	input := "1+2"
	buf := bufio.NewReader(os.Stdin)
	fmt.Fscan(buf, &input)
	//re := regexp.MustCompile(`([0-10IVX]+)(\+/\*){1}([0-10IVX]+)`)
	//re := regexp.MustCompile(`([0-9IVX]+)([\+\*/]{1})([0-9IVX])`)
	//exampleNew := re.FindStringSubmatch(example)
	ex := example.Example{StrExample: input}
	ex.Construct()
	fmt.Println(ex.Result())

}
