package uniqutils

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

type Options struct {
	FlagC       bool
	FlagD       bool
	FlagU       bool
	FlagF       int
	FlagS       int
	FlagI       bool
	// inFilePath  string
	// outFilePath string
}

func Uniq(rows []string, options Options) ([]string, error) {
	if !isOptionsValid(options) {
		return nil, errors.New("cmd options validation error")
	}
	if len(rows) == 0 { // if rows are empty return empty result
		return rows, nil
	}

	rowEqualsCount := 0
	var result []string
	var previousRow []rune

	if options.FlagI {  // Flag -i. making row lowercase
		previousRow = []rune(strings.ToLower(rows[0])) 
	} else {  // Nothing to change
		previousRow = []rune(rows[0]) 
	}

	for i, row := range rows {
		
		if options.FlagI {  // Flag -i. making row lowercase
			// !!!!convert previousRow to Lowercase!!!!!!
			row = strings.ToLower(row)
		}

		rowRune := []rune(row)  // covert byte string to runes slice
		

		rowRuneModified := skipFieldsInRow(rowRune, options.FlagF) // by default -f equals 0 and function doing nothing in this case
		rowRuneModified = skipRunesInRow(rowRuneModified, options.FlagS)  // same as -f option but for -s flag

		previousRowModified := skipFieldsInRow(previousRow, options.FlagF)  // modifying row that uses for comparison with next rows
		previousRowModified = skipRunesInRow(previousRowModified, options.FlagS)

		if slices.Equal(rowRuneModified, previousRowModified) {
			rowEqualsCount += 1
		} else {  // current and previous rows not equals
			ptrRow := formPointerToResultRow(options, previousRow, rowEqualsCount)  // returns nil when conditions required by flags is not completed
			if ptrRow != nil {
				result = append(result, *ptrRow)
			}

			rowEqualsCount = 1  // in any case any row equals itself
			previousRow = rowRune
		}

		if i == (len(rows) - 1) {  // Last row only
			ptrRow := formPointerToResultRow(options, previousRow, rowEqualsCount)
			if ptrRow != nil {
				result = append(result, *ptrRow)
			}
		}
	}

	return result, nil
}

func isOptionsValid(options Options) bool {
	if (options.FlagC && options.FlagD) || (options.FlagC && options.FlagU) || (options.FlagD && options.FlagU) {
		return false
	}
	return true
}

func skipFieldsInRow(row []rune, numField int) []rune {
	fields := 0
	indexAfterSkippedFields := 0
	for i, r := range row {
		if r == ' ' {  // count fields
			fields++
		}
		if fields >= numField {  // reached or step over required numFields
			indexAfterSkippedFields = i + 1
			break
		}
	}
	return row[indexAfterSkippedFields:]
}

func skipRunesInRow(row []rune, numRune int) []rune {
	i := 0
	for ; i < numRune && i < len(row); i++ {}
	return row[i:]
}

func formPointerToResultRow(options Options, previousRow []rune, rowEqualsCount int) *string {
	if ! (options.FlagC || options.FlagD || options.FlagU) {
		row := string(previousRow)
		return &row
	} else if options.FlagC {
		row := fmt.Sprintf("%d %s\n", rowEqualsCount, string(previousRow))
		return &row
	} else if options.FlagD && rowEqualsCount > 1 {  // Flag -d. Only repeated rows
		row := string(previousRow)
		return &row
	} else if options.FlagU && rowEqualsCount == 1 {  // Flag -u. Only unique rows
		row := string(previousRow)
		return &row
	}
	return nil
}
