package main

import (
	"fmt"
	"calc/calcutils"
)

func main() {
	s := calcutils.NewStack()
	fmt.Println(s.IsEmpty())
	fmt.Println(s.Push(1))
	fmt.Println(s.Top())
	fmt.Println(s.Push("abc"))
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.IsEmpty())
}
