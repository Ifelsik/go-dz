package calc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalc(t *testing.T) {
	var testsOk = []struct{
		input    string
		expected float64
	}{
		{
			input: "2+3*4",
			expected: 14,
		},
		{
			input: "2 +3* 4",
			expected: 14,
		},
		{
			input: "3/2+3*4",
			expected: 13.5,
		},
		{
			input: "1 + (2+ (3 +4)*5 +6)+ 7",
			expected: 51,
		},
		{
			input: "-10 + 3",
			expected: -7,
		},
		{
			input: "-(-11-(1*20/2)-11/2*3)",
			expected: 37.5,
		},
	}

	for i, test := range testsOk {
		t.Run(fmt.Sprintf("test %d: ", i), func(t *testing.T) {
			got, err := Calc(test.input)
			require.Equal(t, nil, err,
				"Test %q. Got an error: %v", test.input, err)
			require.Equal(t, test.expected, got,
				"Test %q returned %f; expected %f", test.input, got, test.expected)
		})
	}
}

func TestCalcFail(t *testing.T) {
	var testsFail = map[string] struct {
		input    string
		expected string
	} {
		"wrong brackets balance": {
			input:    "((2+3)",
			expected: "invalid expression",
		},
		"incorrect operation order": {
			input:    "2+*3",
			expected: "something went wrong while pop from stack (probably incorrect expression)",
		},
		"expression contains not a number": {
			input:    "a+2",
			expected: "invalid expression",
		}, 
		"expression contains unknown operator": {
			input:    "4~2",
			expected: "invalid expression",
		},
	}

	for name, test := range testsFail {
		t.Run(name, func(t *testing.T) {
			_, err := Calc(test.input)
			require.EqualError(t, err, test.expected,
				"Test %q expected error: %q; got %q", name, test.expected, err)
		})
	} 
}
