//-------------------------------------------------
// Operators in Go
//-------------------------------------------------

package main

import "fmt"

func main() {

    //-------------------------------------------------
    // OPERATORS
    //-------------------------------------------------
    // An operator is a symbol in a programming language that performs
    // an operation on values, known as operands.
    //
    // For example, in the expression:  10 + 5
    // The '+' is the operator, and 10 and 5 are operands.
    //
    // In Go, operators are used to perform arithmetic, comparison,
    // logical, and assignment operations, among others.

    //-------------------------------------------------
    // ARITHMETIC OPERATORS
    //-------------------------------------------------
    // +   addition
    // -   subtraction
    // *   multiplication
    // /   division
    // %   modulus (remainder)
    //
    // There is no power operator in Go.
    // Use math.Pow(a, b) for exponentiation.

    a, b := 10, 5.5

    fmt.Println(a + 5)   // => 15
    fmt.Println(3.1 - b) // => -2.4
    fmt.Println(a * a)   // => 100
    fmt.Println(a / a)   // => 1
    fmt.Println(11 / 5)  // => 2

	
    // Go is a strongly typed language — type conversions are explicit.
    fmt.Println(a * int(b))     // => 50
    fmt.Println(float64(a) * b) // => 55

    //-------------------------------------------------
    // INC/DEC STATEMENTS
    //-------------------------------------------------
    // The "++" and "--" operators increment or decrement
    // their operands by 1. They are statements, not expressions.
    // Example:
    //    x++ is equivalent to x = x + 1

    x := 10
    x++ // x becomes 11
    x-- // x becomes 10

    //-------------------------------------------------
    // ASSIGNMENT OPERATORS
    //-------------------------------------------------
    //  =   simple assignment
    //  +=  addition assignment
    //  -=  subtraction assignment
    //  *=  multiplication assignment
    //  /=  division assignment
    //  %=  modulus assignment

    a = 10
    a += 2 // => a = 12
    a -= 3 // => a = 9
    a *= 2 // => a = 18
    a /= 3 // => a = 6
    a %= 5 // => a = 1

    //-------------------------------------------------
    // COMPARISON OPERATORS
    //-------------------------------------------------
    // Comparison operators compare two operands and yield
    // a boolean value — either true or false.
    //
    //  ==   equal to
    //  !=   not equal to
    //  >    greater than
    //  <    less than
    //  >=   greater than or equal to
    //  <=   less than or equal to

    fmt.Println(5 == 6)   // => false
    fmt.Println(5 != 6)   // => true
    fmt.Println(10 > 10)  // => false
    fmt.Println(10 >= 10) // => true
    fmt.Println(5 < 5)    // => false
    fmt.Println(5 <= 5)   // => true
	ballance1 := 8
	ballance2 := 9
	fmt.Println(!(ballance1==ballance2 || ballance1!=8))
	fmt.Println(!(ballance1==ballance2 && ballance1!=8))


    //-------------------------------------------------
    // LOGICAL OPERATORS
    //-------------------------------------------------
    // Logical operators are used to combine boolean expressions.
    //
    //  &&   logical AND
    //  ||   logical OR
    //  !    logical NOT

    fmt.Println(0 < 2 && 4 > 1) // => true
    fmt.Println(1 > 5 || 4 > 5) // => false
    fmt.Println(!(1 > 2))       // => true
}
