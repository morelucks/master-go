

// ---------------------------------------------------------
// For Loops in Go
 // ---------------------------------------------------------
/*
A `for` statement is used to execute a block of code repeatedly.
One of the reasons a computer is so useful is that it can repeat operations
many times very quickly — the `for` loop makes this possible in Go.

Structure of a For Loop in Go:
    for initialization; condition; post {
        // block of code to execute
    }

- The **initialization statement** runs once before the loop starts.
- The **condition** is a boolean expression that determines whether the loop continues.
- The **post statement** (like i++) runs after each iteration.

Go’s `for` loop is flexible:
    - It can replace a `while` loop (by omitting initialization and post).
    - It can even create an **infinite loop** (by omitting all three parts).
*/

package main

import "fmt"

func main() {

    // ---------------------------------------------------------
    // 1. Basic For Loop
    // ---------------------------------------------------------
    // printing numbers from 0 to 9
    for i := 0; i < 10; i++ {
        fmt.Println(i)
    }

    // ---------------------------------------------------------
    // 2. For Loop Acting as a While Loop
    // ---------------------------------------------------------
    // Go doesn’t have a while loop — use for instead
    j := 10
    for j >= 0 {
        fmt.Println(j)
        j--
    }

    // ---------------------------------------------------------
    // 3. Multiple Variables in a For Loop
    // ---------------------------------------------------------
    for i, j := 0, 100; i < 10; i, j = i+1, j+1 {
        fmt.Printf("i = %v, j = %v\n", i, j)
    }

    // ---------------------------------------------------------
    // 4. CONTINUE Statement
    // ---------------------------------------------------------
    // Skips the current iteration and moves to the next one
    fmt.Println("\nEven numbers up to 10:")
    for i := 1; i <= 10; i++ {
        if i%2 != 0 {
            continue
        }
        fmt.Println(i)
    }

    // ---------------------------------------------------------
    // 5. BREAK Statement
    // ---------------------------------------------------------
    // Terminates the current loop when a condition is met
    fmt.Println("\nFinding 10 numbers divisible by 13:")
    count := 0
    for i := 0; true; i++ {
        if i%13 == 0 {
            fmt.Printf("%d is divisible by 13\n", i)
            count++
        }

        if count == 10 {
            break
        }
    }

    fmt.Println("Just a message after the for loop.")
}
