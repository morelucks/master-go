package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {

	// ---------------------------------------------------------
	// 1. Unsigned Integer Overflow Example
	// ---------------------------------------------------------
	var x uint8 = 255
	x++
	fmt.Println("Unsigned integer overflow:")
	fmt.Println("x =", x) // Output: 0 (wraps around)
	fmt.Println()

	// ---------------------------------------------------------
	// 2. Compile-Time Overflow Error Example
	// ---------------------------------------------------------
	// Uncomment the following line to see compile-time overflow:
	// var a int8 = 255 + 1 // Error: constant 256 overflows int8
	fmt.Println("Compile-time overflow error example (uncomment to test).")
	fmt.Println()

	// ---------------------------------------------------------
	// 3. Signed Integer Overflow at Runtime
	// ---------------------------------------------------------
	var b int8 = 127
	fmt.Println("Signed integer overflow:")
	fmt.Printf("b + 1 = %d\n", b+1) // Output: -128
	fmt.Println()

	// ---------------------------------------------------------
	// 4. Signed Integer Underflow Example
	// ---------------------------------------------------------
	b = -128
	b--
	fmt.Println("Signed integer underflow:")
	fmt.Println("b =", b) // Output: 127
	fmt.Println()

	// ---------------------------------------------------------
	// 5. Floating-Point Overflow Example
	// ---------------------------------------------------------
	f := float32(math.MaxFloat32)
	fmt.Println("Floating-point overflow:")
	fmt.Println("Initial f =", f)
	f = f * 1.2
	fmt.Println("After overflow f =", f) // Output: +Inf
	fmt.Println()

	// ---------------------------------------------------------
	// 6. Constant Overflow Error
	// ---------------------------------------------------------
	// Uncomment to see compile-time overflow:
	// const i int8 = 300 // Error: 300 overflows int8
	fmt.Println("Constant overflow error example (uncomment to test).")
	fmt.Println()

	// ---------------------------------------------------------
	// 7. Using math/big for Very Large Numbers
	// ---------------------------------------------------------
	fmt.Println("Using math/big for very large numbers:")
	bigNum := new(big.Int)
	bigNum.SetString("999999999999999999999999999999", 10)
	fmt.Println("bigNum =", bigNum)
}
