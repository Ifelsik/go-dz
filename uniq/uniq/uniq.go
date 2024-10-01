package uniq

import (
	"fmt"
	"slices"
	"unicode"
)

type Options struct {
	FlagC       bool
	FlagD       bool
	FlagU       bool
	FlagF       int
	FlagS       int
	FlagI       bool
	InFilePath  string
	OutFilePath string
}

func Uniq(rows []string, options *Options) ([]string, error) {
	if ok, err := IsFlagsValid(options); !ok {
		return nil, fmt.Errorf("options validation error: %v", err)
	}
	if len(rows) == 0 { // if rows are empty return empty result
		return rows, nil
	}

	rowEqualsCount := 0
	var result []string
	var previousRow []rune = []rune(rows[0]) // no index out of range. Case checked upper

	for i, row := range rows {
		rowRune := []rune(row) // covert byte string to runes slice

		// by default -f equals 0 and function doing nothing in this case
		// same as -f option but for -s flag
		// if flag -i providen make rowRune's runes same case
		rowRuneModified := modifyRow(rowRune, options)

		// modifying row that uses for comparison with next rows
		previousRowModified := modifyRow(previousRow, options)

		if slices.Equal(rowRuneModified, previousRowModified) {
			rowEqualsCount += 1
		} else { // current and previous rows not equals
			if resultRow, ok := formResultRow(options, previousRow, rowEqualsCount); ok {
				result = append(result, resultRow)
			}

			rowEqualsCount = 1 // in any case any row equals itself
			previousRow = rowRune
		}

		if i == (len(rows) - 1) { // Last row only
			if resultRow, ok := formResultRow(options, previousRow, rowEqualsCount); ok {
				result = append(result, resultRow)
			}
		}
	}

	return result, nil
}

func IsFlagsValid(options *Options) (bool, error) {
	if (options.FlagC && options.FlagD) || (options.FlagC && options.FlagU) || (options.FlagD && options.FlagU) {
		return false, fmt.Errorf("-c, -d or -u flags can't be used toghter")
	}
	if options.FlagF < 0 {
		return false, fmt.Errorf("flag -f can't be negative")
	}
	if options.FlagS < 0 {
		return false, fmt.Errorf("flag -s can't be negative")
	}
	return true, nil
}

func modifyRow(row []rune, options *Options) []rune {
	// processing fields
	if options.FlagF > 0 {	
		indexAfterSkippedFields := 0
		
		fields := 0
		var rPrev rune // previous rune
		for i, r := range row {
			if (rPrev == ' ' || rPrev == 0) && r != ' ' && r != 0 { // count fields
				fields++
			}
			if fields > options.FlagF { // reached or step over required numFields
				indexAfterSkippedFields = i
				break
			}
			rPrev = r
		}

		if options.FlagF > fields {
			indexAfterSkippedFields = len(row)
		}
		
		row = row[indexAfterSkippedFields:]
	}

	// processing symbols
	if options.FlagS > 0 {
		skippedRunes := options.FlagS
		if options.FlagS > len(row) {
			skippedRunes = len(row)
		}
		row = row[skippedRunes:]
	}

	// convert row to the same case
	if options.FlagI {
		result := make([]rune, len(row))
		for i, r := range row {
			result[i] = unicode.ToLower(r)
		}
		row = result
	}

	return row
}

func formResultRow(options *Options, previousRow []rune, rowEqualsCount int) (string, bool) {
	switch {
	case !(options.FlagC || options.FlagD || options.FlagU):
		row := string(previousRow)
		return row, true
	case options.FlagC:
		row := fmt.Sprintf("%d %s", rowEqualsCount, string(previousRow))
		return row, true
	case options.FlagD && rowEqualsCount > 1: // Flag -d. Only repeated rows
		row := string(previousRow)
		return row, true
	case options.FlagU && rowEqualsCount == 1: // Flag -u. Only unique rows
		row := string(previousRow)
		return row, true
	}
	return "", false
}
