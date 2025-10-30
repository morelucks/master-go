// learning/basics/main.go
// ------------------------------------------
// Variables are how your program remembers values.
// A variable is a name for a memory location where
// a value of a specific type is stored.
//
// There are two ways to declare a variable in Go:
// 1. Using the var keyword
// 2. Using the short declaration operator :=
//
// In Go, semicolons are optional except when writing
// multiple statements on the same line.
// ------------------------------------------

package main

import "fmt"

// Demonstrates variable declarations and usage in Go.
func main() {

	// -------------------------------------------------
	// DECLARING VARIABLES
	// -------------------------------------------------
	// Syntax: var variableName type
	var s1 string = "Learning Go!"
	fmt.Println(s1) // Output: Learning Go!

	// -------------------------------------------------
	// TYPE INFERENCE
	// -------------------------------------------------
	// Go automatically infers the type from the assigned value.
	var k int = 6  // explicit type
	var i = 5      // inferred type: int
	var j = 5.6    // inferred type: float64
	fmt.Println("i:", i, "j:", j, "k:", k)

	// -------------------------------------------------
	// UNUSED VARIABLES
	// -------------------------------------------------
	// Every declared variable must be used, or it causes a compile-time error.
	// The blank identifier (_) can be used to ignore values or silence errors.
	var s2 = "Go!"
	_ = s2 // prevents unused variable error

	// -------------------------------------------------
	// MULTIPLE ASSIGNMENTS
	// -------------------------------------------------
	var ii, jj int
	ii, jj = 5, 8 // tuple assignment
	fmt.Println("Before swap:", ii, jj)

	// Swap values
	ii, jj = jj, ii
	fmt.Println("After swap:", ii, jj)

	// -------------------------------------------------
	// SHORT VARIABLE DECLARATION (:=)
	// -------------------------------------------------
	// Works only inside functions (block scope).
	// Cannot be used at package level.
	s3 := "Learning Golang!"
	_ = s3

	// Multiple short declarations
	car, cost := "Toyota", 50000
	fmt.Println(car, cost)

	// Redeclaration using short syntax
	// At least one variable on the left must be NEW.
	var deleted = false
	deleted, file := true, "a.txt"
	_, _ = deleted, file

	// Short declarations can include expressions
	sum := 5 + 2.5
	fmt.Println(sum)

	// -------------------------------------------------
	// MULTI-LINE VARIABLE DECLARATION
	// -------------------------------------------------
	var (
		age       float64
		firstName string
		gender    bool
	)
	_, _, _ = age, firstName, gender

	// Multiple variables with the same type
	var a, b int
	_, _ = a, b
}
