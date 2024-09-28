package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"uniq/uniq"
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseArguments() (*uniq.Options, error) {
	var cFlag = flag.Bool("c", false, "Count the number of row encounters")
	var dFlag = flag.Bool("d", false, "Show only repeated rows")
	var uFlag = flag.Bool("u", false, "Show only those rows that aren't repeated")
	var fFlag = flag.Int(
		"f",
		0,
		`Skip first 'num_fields' fields in a row.
		Field is not an empty set of characters devided by space`)
	var sFlag = flag.Int(
		"s",
		0,
		`Skip first 'num_chars' chars in a row.
		With '-f' flag counts only chars after 'num fields' fields (except space)`)
	var iFlag = flag.Bool("i", false, "Ignore case")

	flag.Parse()

	args := flag.Args()
	if len(args) > 2 {
		return nil, fmt.Errorf(`maximum 2 arguments can be:
		       inFile and outFile, but got %d`, len(args))
	}

	var inFilePath, outFilePath string
	if len(args) >= 1 {
		inFilePath = args[0]
	}
	if len(args) == 2 {
		outFilePath = args[1]
	}
	
	options := &uniq.Options {
		FlagC:       *cFlag,
		FlagD:       *dFlag,
		FlagU:       *uFlag,
		FlagF:       *fFlag,
		FlagS:       *sFlag,
		FlagI:       *iFlag,
		InFilePath:  inFilePath,
		OutFilePath: outFilePath,
	}
	
	if ok, err := uniq.IsOptionsValid(options); !ok {
		return nil, fmt.Errorf("cmd options validation error: %v", err)
	}

	return options, nil
}

func readFile(options *uniq.Options) ([]string, error) {
	var reader io.Reader
	if len(options.InFilePath) > 0 {  // path isn't empty	
		file, err := os.Open(options.InFilePath)

		if err != nil {
			return nil, err
		}

		defer file.Close()
		reader = file
	} else {
		reader = os.Stdin
		fmt.Println("Hint: to stop input enter Ctrl+Z (Windows)")
	}
	scanner := bufio.NewScanner(reader)

	rows := []string{}
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("while reading got %v", err)
	}
	return rows, nil
}

func writeFile(rows []string, options *uniq.Options) error {
	var outStream io.Writer = os.Stdout
	if len(options.OutFilePath) > 0 {
		file, err := os.Create(options.OutFilePath)  // creates file if doesn't exist
		
		if err != nil {
			return err
		}
		
		defer file.Close()
		outStream = file
	}
	writer := bufio.NewWriter(outStream)

	for _, line := range rows {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("while flush: %v", err)
	}

	return nil
}

func main() {
	options, err := parseArguments()
	check(err)

	rows, err := readFile(options)
	check(err)

	result, err := uniq.Uniq(rows, options)
	check(err)

	err = writeFile(result, options)
	check(err)
}
