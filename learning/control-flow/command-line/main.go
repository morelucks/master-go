// ---------------------------------------------------------
// If Simple Statement + Command Line Arguments
// ---------------------------------------------------------

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// ---------------------------------------------------------
	// 1. Command Line Arguments
	// ---------------------------------------------------------
	// os.Args holds command-line arguments as a slice of strings ([]string).
	// Example:
	// go run main.go GoLang 42
	//
	// os.Args[0] = path to the program
	// os.Args[1] = first argument
	// os.Args[2] = second argument
	fmt.Println("os.Args:", os.Args)

	// Check that enough arguments are provided
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <word> <number>")
		return
	}

	word := os.Args[1]
	rawNum := os.Args[2]

	fmt.Println("Program Path:", os.Args[0])
	fmt.Println("Word Argument:", word)
	fmt.Println("Number Argument (string):", rawNum)
	fmt.Println("Total Arguments:", len(os.Args))

	// ---------------------------------------------------------
	// 2. Converting String to Integer (Regular If)
	// ---------------------------------------------------------
	num, err := strconv.Atoi(rawNum)
	if err != nil {
		fmt.Println("Conversion error:", err)
	} else {
		fmt.Println("Converted Integer:", num)
	}

	// ---------------------------------------------------------
	// 3. Using If with a Simple (Short) Statement
	// ---------------------------------------------------------
	// A simple statement allows you to declare and check in one line.
	// The variables declared are scoped to the if block.
	if n, err := strconv.Atoi("100"); err == nil {
		fmt.Println("No error. Converted value is", n)
	} else {
		fmt.Println("Conversion failed:", err)
	}

	fmt.Println("Program finished successfully.")
}
