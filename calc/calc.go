package main

import (
	"calc/calc"
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Fprintln(os.Stderr, "Error: no expression providen")
		fmt.Println(`Usage example: go run calc.go "2+2*3"`)
		return
	}
	if len(args) > 2 {
		fmt.Fprintln(os.Stderr, "Error: Calc requires 1 string as expression but more than 1 argument was providen")
		return
	}

	result, err := calc.Calc(args[1])

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}

	fmt.Println(result)
}
