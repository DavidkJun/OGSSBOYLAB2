package main

import (
	"flag"
	"fmt"
	lab2 "github.com/roman-mazur/architecture-lab-2"
	"io"
	"os"
	"strings"
)

func main() {
	expression := flag.String("e", "", "Prefix expression to convert")
	fileInput := flag.String("f", "", "File containing prefix expression")
	outputFile := flag.String("o", "", "File to write the result")
	flag.Parse()

	if *expression != "" && *fileInput != "" {
		fmt.Fprintln(os.Stderr, "Error: cannot use both -e and -f options")
		os.Exit(1)
	}

	var input io.Reader
	if *expression != "" {
		input = strings.NewReader(*expression)
	} else if *fileInput != "" {
		file, err := os.Open(*fileInput)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	} else {
		fmt.Fprintln(os.Stderr, "Error: no input provided")
		os.Exit(1)
	}

	var output io.Writer
	if *outputFile != "" {
		file, err := os.Create(*outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		output = file
	} else {
		output = os.Stdout
	}

	handler := &lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}

	if err := handler.Compute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
