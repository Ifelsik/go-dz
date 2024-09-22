package main

import (
	"calc/calcutils"
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Fprintln(os.Stderr, "Error: no expression providen")
		return
	}
	if len(args) > 2 {
		fmt.Fprintln(os.Stderr, "Error: more than 1 expression got")
		return
	}

	result, err := calcutils.Calc(args[0])

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}

	fmt.Println(result)
}
