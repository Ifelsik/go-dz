package calc

import (
	"testing"
	"github.com/stretchr/testify/require"
)

var testsOK = map[string] struct {
	input  []string
	expected []string
} {
	"2+3*4": {
		input: []string{"2", "+", "3", "*", "4"},
		expected: []string{"2", "3", "4", "*", "+"},
	},
	"3/2+3*4": {
		input: []string{"3", "/", "2", "+", "3", "*", "4"},
		expected: []string{"3", "2", "/", "3", "4", "*", "+"},
	},
	"2+(3*5-2)": {
		input: []string{"2", "+", "(", "3", "*", "5", "-", "2", ")"},
		expected: []string{"2", "3", "5", "*", "2", "-", "+"},
	},
	"1+(2+(3+4)*5+6)+7": {
		input: []string{"1", "+", "(", "2", "+", "(", "3", "+", "4", ")", "*", "5", "+", "6", ")", "+", "7"},
		expected: []string{"1", "2", "3", "4", "+", "5", "*", "+", "6", "+", "+", "7", "+"},
	},
	"-3+7": {
		input: []string{"-", "3", "+", "7"},
		expected: []string{"0", "3", "-", "7", "+"},
	},
	"-(-(5+3)*3+(-1))": {
		input: []string{"-", "(", "-", "(", "5", "+", "3", ")", "*", "3", "+", "(", "-", "1", ")", ")"},
		expected: []string{"0", "0", "5", "3", "+", "3", "*", "-", "0", "1", "-", "+", "-"},
	},
	
}

func TestConvertToRPNOK(t *testing.T) {
	for name, test := range testsOK {
		t.Run(name, func(t *testing.T) {
			got, err := ConvertToRPN(test.input)
			require.Equal(t, nil, err,
				"Test %q. Got an error: %v", name, err)
			require.Equal(t, test.expected, got,
				"Test %q returned %v; expected %v", name, got, test.expected)
		})
	}
}
