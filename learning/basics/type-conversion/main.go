// ---------------------------------------------------------
// Type and String Conversions in Go
// ---------------------------------------------------------
// This example demonstrates Go’s explicit type conversion rules,



package main

import (
    "fmt"
)

func main() {

    // ---------------------------------------------------------
    // 1. Converting Between int and float
    // ---------------------------------------------------------
    var x = 3   // int type
    var y = 3.2 // float64 type

    // Invalid operation: mismatched types int and float64
    // x = x * y

    // Convert float64 to int explicitly
    x = x * int(y)
    fmt.Println(x) // => 9

    // Convert int to float64 explicitly
    y = float64(x) * y
    fmt.Println(y) // => 28.8

    // Convert result back to int
    x = int(float64(x) * y)
    fmt.Println(x) // => 259

    // ---------------------------------------------------------
    // 2. Converting Between Integer Types
    // ---------------------------------------------------------
    var a int = 5
    var b int64 = 2

    // int and int64 are different types (explicit conversion required)
    // a = b // compile error
    a = int(b) // correct conversion
    _ = a

    // ---------------------------------------------------------
    // 3. Converting Numbers to Strings
    // ---------------------------------------------------------
    s := string(99)            // converts int (Unicode code point)
    fmt.Println(s)             // => "c"
    fmt.Println(string(34234)) // => 薺 (Unicode for 34234)

    // Convert float to string using fmt.Sprintf
    myStr := fmt.Sprintf("%f", 5.12)
    fmt.Println(myStr) // => "5.120000"

    // Convert int to string using fmt.Sprintf
    myStr1 := fmt.Sprintf("%d", 34234)
    fmt.Println(myStr1) // => "34234"

}
