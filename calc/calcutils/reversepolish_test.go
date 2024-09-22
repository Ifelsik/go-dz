package calcutils

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
}

func TestConvertToRPNOK(t *testing.T) {
	for name, test := range testsOK {
		t.Run(name, func(t *testing.T) {
			got, _ := ConvertToRPN(test.input)
			require.Equal(t, test.expected, got,
				"Test %q returned %v; expected %v", name, got, test.expected)
		})
	}
}
