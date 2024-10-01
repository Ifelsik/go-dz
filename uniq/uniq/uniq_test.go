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
	"without options (empty input)": {
		lines: []string{},
		options: Options{},
		result: []string{},
	},
	"without options (empty string)": {
		lines: []string{""},
		options: Options{},
		result: []string{""},
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
		options: Options{FlagS: 3},
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
	"testing fields according GO's stdlib field definition": {
		lines: []string{
			"  foo bar  baz   qux",
			"a b c qux",
		},
		options: Options{FlagF: 3},
		result: []string{
			"  foo bar  baz   qux",
		},
	},
	"-f 10": {
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
		options: Options{FlagF: 10},
		result: []string{
			"I LOVE MUSIC.",
		},
	},
	"-s 200": {
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
		options: Options{FlagS: 200},
		result: []string{
			"I LOVE MUSIC.",
		},
	},
	"-f 2 -s 50": {
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
		options: Options{FlagS: 200},
		result: []string{
			"I LOVE MUSIC.",
		},
	},
}

func TestUniqOk(t *testing.T) {
	t.Parallel()
	for name, test := range testsOK {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, _ := Uniq(test.lines, &test.options)
			require.Equal(t, test.result, got,
				"Test %q returned %v; expected %v", name, got, test.result) 
		})
	}
}

func TestUniqFail(t *testing.T) {
	t.Parallel()
	t.Run("invalid options", func(t *testing.T) {
		options := Options{FlagC: true, FlagD: true}
		_, err := Uniq([]string{}, &options)
		require.EqualError(t, err,
			"options validation error: -c, -d or -u flags can't be used toghter")
	})
	t.Run("negative -f", func(t *testing.T) {
		t.Parallel()
		options := Options{FlagF: -1}
		_, err := Uniq([]string{}, &options)
		require.EqualError(t, err,
			"options validation error: flag -f can't be negative")
	})
	t.Run("negative -s", func(t *testing.T) {
		t.Parallel()
		options := Options{FlagS: -50}
		_, err := Uniq([]string{}, &options)
		require.EqualError(t, err,
			"options validation error: flag -s can't be negative")
	})
}
