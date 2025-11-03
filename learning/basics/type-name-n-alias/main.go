package main

import "fmt"

func main() {

    // ---------------------------------------------------------
    // 1. TYPE ALIASES
    // ---------------------------------------------------------
    // A type alias gives another name to an existing type.
    // Both names refer to the SAME underlying type.
    // Example: 'byte' is an alias for 'uint8'.
    // ---------------------------------------------------------

    var a uint8 = 10
    var b byte // byte is an alias for uint8

    // Since they share the same type, no conversion is needed
    b = a
    fmt.Printf("a type: %T, b type: %T\n", a, b)

    // Creating an alias 'second' for uint
    type second = uint
    var hour second = 3600

    fmt.Printf("hour type: %T\n", hour) // => hour type: uint
    fmt.Printf("Minutes in an hour: %d\n\n", hour/60) // => 60


    // ---------------------------------------------------------
    // 2. NAMED (DEFINED) TYPES
    // ---------------------------------------------------------
    // A defined type creates a completely NEW type
    // even if it has the same underlying type as another.
    // Conversion is required between the new type and its base type.
    // ---------------------------------------------------------

    type age int        // new type, underlying type: int
    type oldAge age     // new type, underlying type: int (not 'age')
    type veryOldAge age // new type, underlying type: int

    // Defining a new type  with underlying type uint
    type speed uint

    var s1 speed = 10
    var s2 speed = 20

    // Same defined type, arithmetic works
    fmt.Println("Speed difference:", s2-s1)

    var x uint
    // x = s1 // error: different types (uint vs speed)

    //  Must convert manually
    x = uint(s1)
    fmt.Println("Converted speed to uint:", x)

    // Convert back
    var s3 speed = speed(x)
    fmt.Println("Converted uint back to speed:", s3)
}
