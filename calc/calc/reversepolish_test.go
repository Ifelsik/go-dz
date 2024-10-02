package calc

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestConvertToRPNOk(t *testing.T) {
	t.Parallel()
	var testsOk = map[string] struct {
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

	for name, test := range testsOk {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := ConvertToRPN(test.input)
			require.Equal(t, nil, err,
				"Test %q. Got an error: %v", name, err)
			require.Equal(t, test.expected, got,
				"Test %q returned %v; expected %v", name, got, test.expected)
		})
	}
}

func TestConvertToRPNFail(t *testing.T) {
	t.Parallel()
	var testsFail = map[string] struct {
		input    []string
		expected string
	} {
		"invalid operator": {
			input: []string{"2", "+", "(", "3", "**", "2", ")"},
			expected: "unknown operator",
		},
	}

	for name, test := range testsFail {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			_, err := ConvertToRPN(test.input)
			require.EqualError(t, err, test.expected,
				"Test %q expected err %q; got err %q", name, test.expected, err)
		})
	}
}
