package calc

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertToRPN(expression []string) ([]string, error) {
	result := []string{}
	stack := NewStack()
	for _, str := range expression {
		if _, err:= strconv.ParseFloat(str, 64); err == nil {  // got a number, just add it to result
			result = append(result, str)
			continue
		}
		
		if !stack.IsEmpty() {
			top, _ := stack.Top()  // error checked in IsEmpty method!!!!

			currentPriority, _ := getOperandPriority(str)
			topPriority, _ 	   := getOperandPriority(top)

			isBracketsEnd := str == ")"
			
			for isBracketsEnd || (topPriority <= currentPriority && top != "(" ) {  // top priority equals or higher than current
				result = append(result, top)

				stack.Pop()

				var err error
				top, err = stack.Top()

				if err != nil {  // stack is empty
					break
				}

				topPriority, _ = getOperandPriority(top)

				if top == "(" {
					stack.Pop()
					isBracketsEnd = false
				}
			}
		}

		if str != ")" {
			stack.Push(str)
		}
	}

	for !stack.IsEmpty() {
		elem, _ := stack.Pop()
		result = append(result, elem)
	}

	return result, nil
}


// getOperandPriority returns priority of given operand
// It accepets "+", "-", "*", "/", "(", ")"
// Returns for given operand integer priority (0 - highest, 1 - lower, 2 - much more lower etc.)
func getOperandPriority(op string) (int8, error) {
	var priority int8 = -1
	var err error
	op = strings.TrimSpace(op)  // remove sapces from both sides of string
	switch op {
		case "+":
			priority = 3
		case "-":
			priority = 3
		case "*":
			priority = 2
		case "/":
			priority = 2
		case "(":
			priority = 1
		case ")":
			priority = 1
		default:
			err = fmt.Errorf("unknown operator") 
	}

	return priority, err
}

