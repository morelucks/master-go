// ---------------------------------------------------------
// ARRAYS IN GO
// ---------------------------------------------------------
// Arrays are fixed-length sequences of elements of the same type.
// Once declared, their size cannot be changed.
// Arrays are value types — assigning one array to another copies all elements.
// ---------------------------------------------------------

package main

import "fmt"

func main() {

    // ---------------------------------------------------------
    // 1. Declaring an Array
    // ---------------------------------------------------------
    var numbers [4]int

    // Array zero value is the zeroed value of its element type.
    fmt.Printf("%v\n", numbers)  // -> [0 0 0 0]
    fmt.Printf("%#v\n", numbers) // -> [4]int{0, 0, 0, 0}

    // ---------------------------------------------------------
    // 2. Array Initialization
    // ---------------------------------------------------------
    var a1 = [4]float64{}                           // initialized with default zero values
    var a2 = [3]int{5, -3, 5}                       // initialized with given values
    a3 := [4]string{"Dan", "Diana", "Paul", "John"} // short declaration
    a4 := [4]string{"x", "y"}                       // partially initialized

    // The ellipsis operator (...) automatically infers the array length.
    a5 := [...]int{1, 4, 5}
    fmt.Println("The length of a5 is:", len(a5)) // len is 3

    // Declaring array across multiple lines (comma required)
    a6 := [...]int{
        1,
        2,
        3,
    }

    // avoid “declared and not used” errors
    _, _, _, _, _, _ = a1, a2, a3, a4, a5, a6

    // ---------------------------------------------------------
    // 3. Modifying Arrays
    // ---------------------------------------------------------
    // Arrays are fixed in length — cannot add/remove elements
    numbers[0] = 7              // modifying first element (index 0)
    fmt.Printf("%v\n", numbers) // -> [7 0 0 0]

    // Invalid (out of bounds)
    // numbers[5] = 8

    // Accessing array elements
    x := numbers[0]
    fmt.Println("x is", x) // => x is 7

    // ---------------------------------------------------------
    // 4. Iterating Through Arrays
    // ---------------------------------------------------------
    fmt.Println("\nIterating using range:")
    for i, v := range numbers {
        fmt.Println("index:", i, "value:", v)
    }

    fmt.Println("\nIterating using classic for loop:")
    for i := 0; i < len(numbers); i++ {
        fmt.Println("index:", i, "value:", numbers[i])
    }

    // ---------------------------------------------------------
    // 5. Multi-Dimensional Arrays
    // ---------------------------------------------------------
    balances := [2][3]int{
        [3]int{5, 6, 7}, // explicit type
        {8, 9, 10},      // implicit type
    }

    fmt.Println("\nMulti-dimensional array iteration:")
    for _, arr := range balances {
        for _, value := range arr {
            fmt.Printf("%d ", value)
        }
        fmt.Println("")
    }

    // ---------------------------------------------------------
    // 6. Copying Arrays
    // ---------------------------------------------------------
    // Arrays are copied by value — changes to one do not affect the other
    m := [3]int{1, 2, 3}
    n := m // copy of m

    fmt.Println("\nn is equal to m:", n == m) // true
    m[0] = -1                                 // only m is modified
    fmt.Println("n is equal to m:", n == m)   // false
}
