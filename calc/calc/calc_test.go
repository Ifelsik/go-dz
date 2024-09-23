package calc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

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

func TestCalc(t *testing.T) {
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
