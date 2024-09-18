package uniqutils

import (
	"testing"
	"github.com/stretchr/testify/require"
)

var testsOK = map[string] struct {
	lines   []string
	options Options
	result  []string
}{
	"without options #1": {
		lines: []string{
			"abc",
			"abc",
			"abc",
		},
		options: Options{},
		result: []string{
			"abc",
		},
	},
	"without options #2": {
		lines: []string{
			"abc",
			"def",
			"def",
			"",
			"abc",
			"def",
			"def",
			"ghi",
		},
		options: Options{},
		result: []string{
			"abc",
			"def",
			"",
			"abc",
			"def",
			"ghi",
		},
	},
	"without options #3 (empty input)": {
		lines: []string{},
		options: Options{},
		result: []string{},
	},
	"without options #4 (empty string)": {
		lines: []string{""},
		options: Options{},
		result: []string{""},
	},
	"with -c flag": {
		lines: []string{
			"abc",
			"def",
			"def",
			"",
			"abc",
			"def",
			"def",
			"ghi",
		},
		options: Options{FlagC: true},
		result: []string{
			"1 abc",
			"2 def",
			"1 ",
			"1 abc",
			"2 def",
			"1 ghi",
		},
	},
	"with -d flag": {
		lines: []string{
			"abc",
			"def",
			"def",
			"",
			"abc",
			"def",
			"def",
			"ghi",
		},
		options: Options{FlagD: true},
		result: []string{
			"def",
			"def",
		},
	},
	"with -u flag": {
		lines: []string{
			"abc",
			"def",
			"def",
			"",
			"abc",
			"def",
			"def",
			"ghi",
		},
		options: Options{FlagU: true},
		result: []string{
			"abc",
			"",
			"abc",
			"ghi",
		},
	},
	"with -f flag": {
		lines: []string{
			"ab ab",
			"cd ab",
			"ef ab",
			"",
			"ab",
			"ab cd",
			"ab cd ef",
		},
		options: Options{FlagF: 1},
		result: []string{
			"ab ab",
			"",
			"ab",
			"ab cd",
			"ab cd ef",
		},
	},
	// "with -f flag (overflow)": {
	// },
	"with -f and -c flags": {
		lines: []string{
			"ab ab",
			"cd ab",
			"ef ab",
			"",
			"ab",
			"ab",
			"ab cd",
			"ab cd ef",
		},
		options: Options{FlagF: 1, FlagC: true},
		result: []string{
			"3 ab ab",
			"1 ",
			"2 ab",
			"1 ab cd",
			"1 ab cd ef",
		},
	},
	"with -s flag": {
		lines: []string{
			"abcd",
			"cdcd",
			"",
			"a",
			"aa",
			"ffff",
			"ffff",
		},
		options: Options{FlagS: 2},
		result: []string{
			"abcd",
			"",
			"ffff",
		},
	},
	"with  -f, -s and -c flag": {
		lines: []string{
			"ab a ab",
			"cd b ab",
			"ef c ab",
			"",
			"ab",
			"ab ab",
			"abab",
		},
		options: Options{FlagF: 1, FlagS: 1, FlagC: true},
		result: []string{
			"3 ab a ab",
			"1 ",
			"2 ab",
			"1 abab",
		},
	},
	"with flag -i": {
		lines: []string{
			"AbC",
			"aBc",
			"",
			"abcd",
			"abc",
			"DDDD",
			"dddd",
		},
		options: Options{FlagI: true},
		result: []string{
			"AbC",
			"",
			"abcd",
			"abc",
			"DDDD",
		},
	},
}

func TestUniqOk(t *testing.T) {
	for name, test := range testsOK {
		t.Run(name, func(t *testing.T) {
			// t.Parallel()
			got, _ := Uniq(test.lines, test.options)
			require.Equal(t, test.result, got,
				          "Test %q returned %v; expected %v", name, got, test.result) 

		})
	}
}

// func TestUniqFail(t *testing.T) {
// 	for name, test := range testsFail {

// 	}
// }
