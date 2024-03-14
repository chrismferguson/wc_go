package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main()  {
	// define flags for the different command options
	countBytes := flag.Bool("c", false, "count bytes")
	countLines := flag.Bool("l", false, "count lines")
	countWords := flag.Bool("w", false, "count words")

	// parse the flags
	flag.Parse()

	// function to count from an io.Reader
	countFromReader := func(r io.Reader) (lines, words, bytes int){
		scanner := bufio.NewScanner(r)
		for scanner.Scan(){
			bytes += len(scanner.Bytes()) // count bytes
			lines++
			words += len(strings.Fields(scanner.Text())) // count words
		}
		return
	}

	// check if input is from a file or standard input
	var input io.Reader
	if flag.NArg() > 0{
		fileName := flag.Arg(0)
		file, err := os.Open(fileName)
		if err != nil{
			fmt.Println("Error opening file", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	} else {
		input = os.Stdin
	}

	// perform counting
	lines, words, bytes := countFromReader(input)

	// determine what to print based on flags 
	printCounts := func() {
		if *countLines{
			fmt.Printf("%8d ", lines)
		}
		if *countWords{
			fmt.Printf("%8d ", words)
		}
		if *countBytes{
			fmt.Printf("%8d ", bytes)
		}
		if !*countLines && !*countWords && !*countBytes {
			// if no flags are provided, print all counts
			fmt.Printf("%8d%8d%8d ", lines, words, bytes)
		}
		fmt.Println()
	}

	// print the counts
	printCounts()
}
