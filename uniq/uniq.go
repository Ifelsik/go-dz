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

func main () {
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
		fmt.Printf(`Error: maximum 2 arguments can be:
					inputFile and outFile, but got %d`, len(args))
		return
	}
	var inFilePath, outFilePath string
	if len(args) >= 1 {
		inFilePath = args[0]
	}
	if len(args) == 2 {
		outFilePath = args[1]
	}
	
	options := uniq.Options {
		FlagC: *cFlag,
		FlagD: *dFlag,
		FlagU: *uFlag,
		FlagF: *fFlag,
		FlagS: *sFlag,
		FlagI: *iFlag,
	}

	var reader io.Reader
	if len(inFilePath) > 0 {  // path isn't empty	
		file, err := os.Open(inFilePath)
		check(err)
		defer file.Close()
		reader = file
	} else {
		reader = os.Stdin
		fmt.Println("Hint: To stop the input press Ctrl + C")
	}
	scanner := bufio.NewScanner(reader)

	rows := []string{}
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error: while reading ", err)
	} 

	result, err := uniq.Uniq(rows, options)
	if err != nil {
		fmt.Printf("Got an error: %s", err)
	}
	
	var outStream io.Writer
	if len(outFilePath) > 0 {
		file, err := os.Create(outFilePath)  // creates file if doesn't exist
		check(err)
		defer file.Close()
		outStream = file
	} else {
		outStream = os.Stdout
	}
	writer := bufio.NewWriter(outStream)

	for _, line := range result {
		_, err := writer.WriteString(line)
		check(err)
	}
}
