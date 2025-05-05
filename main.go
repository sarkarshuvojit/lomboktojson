package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/sarkarshuvojit/lomboktojson/pkg"
)



func main() {
	// Define command-line flags
	inputFile := flag.String("i", "", "Input file (optional)")
	flag.Parse()

	var input string
	// Check if there's data being piped in
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// Data is being piped in
		reader := bufio.NewReader(os.Stdin)
		bytes, err := io.ReadAll(reader)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
			os.Exit(1)
		}
		input = string(bytes)
	} else if *inputFile != "" {
		// Read from input file
		bytes, err := os.ReadFile(*inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
			os.Exit(1)
		}
		input = string(bytes)
	} else {
		// Read from command line arguments
		if len(flag.Args()) > 0 {
			input = flag.Args()[0]
		} else {
			// No input provided - prompt user
			fmt.Println("Enter input text (press Ctrl+D when done):")
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				input += scanner.Text() + "\n"
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
				os.Exit(1)
			}
		}
	}

	// Process the input
	output, err := pkg.LombokToJson(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
	fmt.Print(output)
}
