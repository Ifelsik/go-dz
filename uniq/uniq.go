package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"uniq/uniqutils"
)

// debug mode enabled
const DEBUG = true

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
	// if len(args) > 2 {
	// 	fmt.Errorf("> 2 args")
	// }
	var inFilePath, outFilePath string
	if len(args) >= 1 {
		inFilePath = args[0]
	}
	if len(args) == 2 {
		outFilePath = args[1]
	}
	
	options := uniqutils.Options {
		FlagC: *cFlag,
		FlagD: *dFlag,
		FlagU: *uFlag,
		FlagF: *fFlag,
		FlagS: *sFlag,
		FlagI: *iFlag,
	}

	var reader io.Reader
	if len(inFilePath) > 0 {  // path isn't empty
		var err error	
		reader, err = os.Open(inFilePath)
		check(err)
		// defer reader.Close()
	} else {
		reader = os.Stdin
		fmt.Println("Hint: To stop the input press Ctrl + C")
	}
	scanner := bufio.NewScanner(reader)

	rows := []string{}
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}
	// check errors

	if DEBUG {
		fmt.Println("-c: ", *cFlag)
		fmt.Println("-d: ", *dFlag)
		fmt.Println("-u: ", *uFlag)
		fmt.Println("-f: ", *fFlag)
		fmt.Println("-s: ", *sFlag)
		fmt.Println("-i: ", *iFlag)

		fmt.Println("file in: ", inFilePath)
		fmt.Println("file out:", outFilePath)
	}

	result, _ := uniqutils.Uniq(rows, options) 
	
	var writer io.Writer
	if len(outFilePath) > 0 {
		writer, err := os.Create(outFilePath)  // creates file if doesn't exist
		check(err)
		defer writer.Close()
	} else {
		writer = os.Stdout
	}
	writer_ := bufio.NewWriter(writer) // rename!!!

	for _, line := range result {
		_, err := writer_.WriteString(line)
		check(err)
	}
	

	if DEBUG {
		fmt.Println(rows)
		fmt.Println(result)
	}
}
