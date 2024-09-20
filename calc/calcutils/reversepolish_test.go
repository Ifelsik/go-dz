package calcutils

import (
	"testing"
	"github.com/stretchr/testify/require"
)

var testsOK = map[string] struct {
	input  []string
	result []string
}{
	"2 + 3": {
		input: []string{"2", "+", "3"},
		result: []string{"2", "3", "+"},
	},
	"2 + 3 * 4": {
		input: []string{"2", "+", "3", "*", "4"},
		result: []string{"2", "3", "4", "*", "+"},
	},
	"3 / 2 + 3 * 4": {
		input: []string{"3", "/", "2", "+", "3", "*", "4"},
		result: []string{"3", "2", "/", "3", "4", "*", "+"},
	},
}

func TestConvertToRPNOK(t *testing.T) {
	for name, test := range testsOK {
		t.Run(name, func(t *testing.T) {
			got, _ := ConvertToRPN(test.input)
			require.Equal(t, test.result, got,
				"Test %q returned %v; expected %v", name, got, test.result)
		})
	}
}
