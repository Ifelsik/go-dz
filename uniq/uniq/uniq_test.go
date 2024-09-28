package uniq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testsOK = map[string] struct {
	lines   []string
	options Options
	result  []string
}{	
	"examples/without params": {
		lines: []string{
			"I love music.",
			"I love music.",
			"I love music.",
			"",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I love music of Kartik.",
		},
		options: Options{},
		result: []string{
			"I love music.",
			"",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
		},
	},
	"examples/-c": {
		lines: []string{
			"I love music.",
			"I love music.",
			"I love music.",
			"",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I love music of Kartik.",
		},
		options: Options{FlagC: true},
		result: []string{
			"3 I love music.",
			"1 ",
			"2 I love music of Kartik.",
			"1 Thanks.",
			"2 I love music of Kartik.",
		},
	},
	"examples/-d": {
		lines: []string{
			"I love music.",
			"I love music.",
			"I love music.",
			"",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I love music of Kartik.",
		},
		options: Options{FlagD: true},
		result: []string{
			"I love music.",
			"I love music of Kartik.",
			"I love music of Kartik.",
		},
	},
	"examples/-u": {
		lines: []string{
			"I love music.",
			"I love music.",
			"I love music.",
			"",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I love music of Kartik.",
		},
		options: Options{FlagU: true},
		result: []string{
			"",
			"Thanks.",
		},
	},
	"examples/-i": {
		lines: []string{
			"I LOVE MUSIC.",
			"I love music.",
			"I LoVe MuSiC.",
			"",
			"I love MuSIC of Kartik.",
			"I love music of kartik.",
			"Thanks.",
			"I love music of kartik.",
			"I love MuSIC of Kartik.",
		},
		options: Options{FlagI: true},
		result: []string{
			"I LOVE MUSIC.",
			"",
			"I love MuSIC of Kartik.",
			"Thanks.",
			"I love music of kartik.",
		},
	},
	"examples/-f_num": {
		lines: []string{
			"We love music.",
			"I love music.",
			"They love music.",
			"",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks.",
		},
		options: Options{FlagF: 1},
		result: []string{
			"We love music.",
			"",
			"I love music of Kartik.",
			"Thanks.",
		},
	},
	"examples/s_num": {
		lines: []string{
			"I love music.",
			"A love music.",
			"C love music.",
			"",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks.",
		},
		options: Options{FlagS: 1},
		result: []string{
			"I love music.",
			"",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks.",
		},
	},
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
			got, _ := Uniq(test.lines, &test.options)
			require.Equal(t, test.result, got,
				          "Test %q returned %v; expected %v", name, got, test.result) 
		})
	}
}


func TestUniqFail(t *testing.T) {
	t.Run("invalid options", func(t *testing.T) {
		options := Options{FlagC: true, FlagD: true}
		_, err := Uniq([]string{}, &options)
		require.EqualError(t, err,
			"options validation error: -c, -d or -u flags can't be used toghter")
	})
	t.Run("negative -f", func(t *testing.T) {
		options := Options{FlagF: -1}
		_, err := Uniq([]string{}, &options)
		require.EqualError(t, err,
			"options validation error: flag -f can't be negative")
	})
	t.Run("negative -s", func(t *testing.T) {
		options := Options{FlagS: -50}
		_, err := Uniq([]string{}, &options)
		require.EqualError(t, err,
			"options validation error: flag -s can't be negative")
	})
}
