
// This example demonstrates how conditional statements work in Go.
// Similar to Solidity's "if", "else if", and "else" logic,
// Go allows strict, boolean-based conditions without truthy or falsy behavior.
//
// Covers:
// 1. Basic if Statement
// 2. if with Logical Operators
// 3. else if and else Branches
// 4. No “Truthiness” in Go (explicit boolean evaluation)
//
// Author: Lucky Kamshak
// Date: 4th November 2025
// ---------------------------------------------------------

package main

import "fmt"

func main() {

    // ---------------------------------------------------------
    // 1. Basic if Statement
    // ---------------------------------------------------------
    // Syntax:
    // if condition_that_evaluates_to_boolean {
    //     perform action
    // }

    price, inStock := 100, true

    if price >= 80 { // parentheses are NOT required in Go
        fmt.Println("Too Expensive")
    }

    // ---------------------------------------------------------
    // 2. if with Logical Operators (&&, ||)
    // ---------------------------------------------------------
    // You can combine multiple conditions using logical operators.
    if price <= 100 && inStock == true {
        fmt.Println("Buy it!")
    }

    // Note: the same logic can be simplified as:
    // if price <= 100 && inStock { }

    // ---------------------------------------------------------
    // 3. No “Truthiness” in Go
    // ---------------------------------------------------------
    // In Go, only expressions that explicitly evaluate to boolean
    // can be used in conditionals.
    //
    // The following would cause a compile error:
    //
    // if price {
    //     fmt.Println("We have price!")
    // }

    // ---------------------------------------------------------
    // 4. else if and else Branches
    // ---------------------------------------------------------
    // Only one branch will be executed in this chain.
    if price < 100 {
        fmt.Println("It's cheap!")
    } else if price == 100 {
        fmt.Println("On the edge")
    } else { // executed only if all previous conditions are false
        fmt.Println("It's Expensive!")
    }

}
