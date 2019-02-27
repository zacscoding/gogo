package main

import (
	"fmt"
	"github.com/zacscoding/gogo/calculator"
)

func main() {
	// fmt.Println(calculator.Add(1, 2))
	fmt.Print("calculator.Eval(\"3+5\") :: ", calculator.Eval("3 + 5"))
}
