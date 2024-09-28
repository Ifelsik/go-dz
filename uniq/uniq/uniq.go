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
	if ok, err := IsOptionsValid(options); !ok {
		return nil, fmt.Errorf("options validation error: %v", err)
	}
	if len(rows) == 0 { // if rows are empty return empty result
		return rows, nil
	}

	rowEqualsCount := 0
	var result []string
	var previousRow []rune = []rune(rows[0])  // no index out of range. Case checked upper


	for i, row := range rows {
		rowRune := []rune(row)  // covert byte string to runes slice
  
		// by default -f equals 0 and function doing nothing in this case
		// same as -f option but for -s flag
		// if flag -i providen make rowRune's runes same case 
		rowRuneModified := modifyRow(rowRune, options)
		
		// modifying row that uses for comparison with next rows
		previousRowModified := modifyRow(previousRow, options)

		if slices.Equal(rowRuneModified, previousRowModified) {
			rowEqualsCount += 1
		} else {  // current and previous rows not equals
			if resultRow, ok := formResultRow(options, previousRow, rowEqualsCount); ok {
				result = append(result, resultRow)
			}

			rowEqualsCount = 1  // in any case any row equals itself
			previousRow = rowRune
		}

		if i == (len(rows) - 1) {  // Last row only
			if resultRow, ok := formResultRow(options, previousRow, rowEqualsCount); ok {
				result = append(result, resultRow)
			}
		}
	}

	return result, nil
}

func IsOptionsValid(options *Options) (bool, error) {
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
	// processing flag -f
	fields := 0
	indexAfterSkippedFields := 0
	for i, r := range row {
		if r == ' ' {  // count fields
			fields++
		}
		if fields >= options.FlagF {  // reached or step over required numFields
			indexAfterSkippedFields = i + 1
			break
		}
	}

	row = row[indexAfterSkippedFields:]

	// processing flag -s
	skippedRunes := 0
	for ; skippedRunes < options.FlagS && skippedRunes < len(row); skippedRunes++ {}

	row = row[skippedRunes:]

	// processing flag -f
	if options.FlagI {
		result := make([]rune, len(row))
		for _, r := range row {
			result = append(result, unicode.ToLower(r))
		}
		row = result
	}

	return row
}

func formResultRow(options *Options, previousRow []rune, rowEqualsCount int) (string, bool) {
	switch {
	case ! (options.FlagC || options.FlagD || options.FlagU):
		row := string(previousRow)
		return row, true
	case options.FlagC:
		row := fmt.Sprintf("%d %s", rowEqualsCount, string(previousRow))
		return row, true
	case options.FlagD && rowEqualsCount > 1: // Flag -d. Only repeated rows
		row := string(previousRow)
		return row, true
	case options.FlagU && rowEqualsCount == 1 :  // Flag -u. Only unique rows
		row := string(previousRow)
		return row, true
	}	
	return "", false
}
