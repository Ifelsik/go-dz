package calc

import (
	"fmt"
	"strconv"
	"strings"
	"regexp"
)

// It accepts expression in a slice format (i.e. "2+3" -> ["2", "+", "3"])
// Returns result and error
func Calc(expression string) (float64, error) {
	if !isValidExpression(expression) {
		return 0, fmt.Errorf("invalid expression")
	}
	// expression = strings.ReplaceAll(expression, " ", "")  // deletes spaces

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

		value1, err1 := stack.Pop()
		value2, err2 := stack.Pop()

		if err1 != nil || err2 != nil {
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
			stack.Push(strconv.FormatFloat(value2Float / value1Float, 'f', -1, 64))
		default:
			return 0, fmt.Errorf("unknown operator")
		}
	}

	result, err := stack.Top()
	if err != nil {
		return 0, err
	}

	resultFloat, err := strconv.ParseFloat(result, 64)
	if err != nil {
		return 0, fmt.Errorf("conversation error")
	}  
	
	return resultFloat, nil
}

func isValidExpression(expression string) bool {
	expression = strings.ReplaceAll(expression, " ", "")  // deletes spaces

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
