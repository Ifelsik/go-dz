package calc

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// It accepts expression in string.
// Returns result and error
func Calc(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, " ", "")  // deletes spaces
	
	if !isValidExpression(expression) {
		return 0, fmt.Errorf("invalid expression")
	}

	re := regexp.MustCompile(`\d+|[+*/()-]`)
	expressionTokenized := re.FindAllString(expression, -1)

	infixExpression, err := ConvertToRPN(expressionTokenized)
	if err != nil {
		return 0, err
	}

	stack := NewStack()

	for _, op := range infixExpression {
		if _, err := strconv.ParseFloat(op, 64); err == nil {
			stack.Push(op)
			continue
		}

		value1, ok1 := stack.Pop()
		value2, ok2 := stack.Pop()

		if !(ok1 && ok2) {
			return 0, fmt.Errorf(
			       "something went wrong while pop from stack (probably incorrect expression)")
		}

		value1Float, err1 := strconv.ParseFloat(value1, 64)
		value2Float, err2 := strconv.ParseFloat(value2, 64)

		if err1 != nil || err2 != nil {
			return 0, fmt.Errorf("something went wrong while atof conversation") 
		}

		switch op {
		case "+":
			stack.Push(strconv.FormatFloat(value2Float + value1Float, 'f', -1, 64))
		case "-":
			stack.Push(strconv.FormatFloat(value2Float - value1Float, 'f', -1, 64))
		case "*":
			stack.Push(strconv.FormatFloat(value2Float * value1Float, 'f', -1, 64))
		case "/":
			if value1Float == 0 {  // precision approx. 1e-300
				return 0, fmt.Errorf("division by zero")
			}
			stack.Push(strconv.FormatFloat(value2Float / value1Float, 'f', -1, 64))
		default:
			return 0, fmt.Errorf("unknown operator")
		}
	}

	result, ok := stack.Top()
	if !ok {
		return 0, fmt.Errorf("after calculation top of stack is empty")
	}

	resultFloat, err := strconv.ParseFloat(result, 64)
	if err != nil {
		return 0, fmt.Errorf("conversation error")
	}  
	
	return resultFloat, nil
}

func isValidExpression(expression string) bool {
	// Acceptable chars
	if matched, _ := regexp.MatchString(`[^0-9+\-*/().]`, expression); matched {
		return false
	}

	// brackets balance
	balance := 0
	for _, char := range expression {
		if char == '(' {
			balance++
		} else if char == ')' {
			balance--
			if balance < 0 {
				return false
			}
		}
	}

	if balance != 0 {
		return false
	}

	return true
}
