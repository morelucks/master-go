package main

import (
	"fmt"
	"math"
)

func main() {
	// -------------------------------------------------
	// The main function
	// -------------------------------------------------
	// 'main' is a special function name in Go.
	// It serves as the entry point for the executable program.
	// Go uses braces { } to delimit code blocks.

	// -------------------------------------------------
	// Local scoped variables and constants
	// -------------------------------------------------
	// Variables declared inside functions are local to that function.
	// Constants are values that cannot change once declared.

	var a int = 7
	var b float64 = 3.5
	const c int = 10

	fmt.Println("Hello Go world!") // => Hello Go world!
	fmt.Println(a, b, c)           // => 7 3.5 10

	// -------------------------------------------------
	// Using math package
	// -------------------------------------------------
	x := 5.2
	powX := math.Pow(x, 2)
	fmt.Println("The power of x is:", powX)



	//x := 5.2
	//xx := math.Pow(x, 2)

	//fmt.Println("the powoer of x is : ", xx)
}