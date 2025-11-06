
// ---------------------------------------------------------
// SWITCH STATEMENT
// ---------------------------------------------------------
// The switch statement provides a cleaner alternative to
// multiple if-else conditions. Behind the scenes, Go converts
// switch statements into if-else chains automatically.
//
// Each case must be comparable, and a default clause can be
// used as a fallback (similar to an else statement).
// ---------------------------------------------------------

package main

import "fmt"

func main() {

    // ---------------------------------------------------------
    // 1. Basic Switch Statement
    // ---------------------------------------------------------
    language := "golang"

    switch language {
    case "Python": // compare string to string
        fmt.Println("You are learning Python! You don't use { } but indentation !!")
    case "Go", "golang": // compare language with "Go" OR "golang"
        fmt.Println("Good, Go for Go! You are using {}!")
    default:
        fmt.Println("Any other programming language is a good start!")
    }

    // ---------------------------------------------------------
    // 2. Switch with Boolean Expressions
    // ---------------------------------------------------------
    n := 5

    // comparing the result of an expression which is bool to another bool value
    switch true {
    case n%2 == 0:
        fmt.Println("Even!")
    case n%2 != 0:
        fmt.Println("Odd!")
    default:
        fmt.Println("Never here!")
    }

    // ---------------------------------------------------------
    // 3. Switch with Simple Statement
    // ---------------------------------------------------------
    // Syntax: statement (n := 10), semicolon, and a switch condition (true)
    // The condition 'true' can be omitted because it’s the default.
    switch n := 10; true {
    case n > 0:
        fmt.Println("Positive")
    case n < 0:
        fmt.Println("Negative")
    default:
        fmt.Println("Zero")
    }
}
