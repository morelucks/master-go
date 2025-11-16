//-------------------------------------------------
// Understanding Slices & Backing Arrays in Go
//-------------------------------------------------
//
// This example demonstrates how Go slices work — including
// slice headers, backing arrays, shared memory behavior,
// append mechanics, capacity growth, and the copy() function.
//
// Ethereum-inspired variable names used for clarity.
//-------------------------------------------------

package main

import "fmt"

func main() {

    //-------------------------------------------------
    // SLICE HEADER & BACKING ARRAY
    //-------------------------------------------------
    //
    // A slice in Go is NOT an array.
    // It is a small structure with 3 fields:
    // 1. Pointer → the backing array
    // 2. Length  → len()
    // 3. Capacity → cap()
    //
    // Slicing DOES NOT create a new array.
    // Multiple slices can share the same backing array.
    //-------------------------------------------------

    s1 := []int{10, 20, 30, 40, 50}
    s3 := s1[0:2] // shares the same backing array
    s4 := s1[1:3] // also shares the same backing array

    s3[1] = 600 // modifies the backing array
    fmt.Println("s1:", s1) // → [10 600 30 40 50]
    fmt.Println("s4:", s4) // → [600 30]

    //-------------------------------------------------
    // SLICING AN ARRAY
    //-------------------------------------------------
    // If you slice an array, the array becomes
    // the backing array of the new slice.
    //-------------------------------------------------

    arr := [4]int{10, 20, 30, 40}
    slice1 := arr[0:2]
    slice2 := arr[1:3]

    arr[1] = 2 // modifies array → affects both slices
    fmt.Println(arr, slice1, slice2)
    // → [10 2 30 40] [10 2] [2 30]

    _, _, _ = slice1, slice2, s4

    //-------------------------------------------------
    // APPEND() AND BACKING ARRAY REALLOCATION
    //-------------------------------------------------
    //
    // append() may or may not allocate a new backing array.
    // If the existing array has enough capacity → reuse.
    // If not → Go allocates a bigger array.
    //-------------------------------------------------

    cars := []string{"Ford", "Honda", "Audi"}
    newCars := []string{}

    // newCars gets its OWN backing array
    newCars = append(newCars, cars[0:2]...)

    cars[0] = "Nissan"
    fmt.Println("cars:", cars)
    fmt.Println("newCars:", newCars)
    // → cars modified ONLY, newCars unchanged

    //-------------------------------------------------
    // APPEND() GROWTH RULES (LENGTH & CAPACITY)
    //-------------------------------------------------
    //
    // Capacity grows automatically:
    // 1 → 2 → 4 → 8 → 16 …
    // This avoids frequent reallocations.
    //-------------------------------------------------

    nums := []int{1}
    fmt.Println(len(nums), cap(nums)) // 1 1

    nums = append(nums, 2)
    fmt.Println(len(nums), cap(nums)) // 2 2

    nums = append(nums, 3)
    fmt.Println(len(nums), cap(nums)) // 3 4

    nums = append(nums, 4, 5)
    fmt.Println(len(nums), cap(nums)) // 5 8

    //-------------------------------------------------
    // COPY() FUNCTION
    //-------------------------------------------------
    //
    // copy(dst, src) copies up to the MIN length
    // and returns the number of elements copied.
    //-------------------------------------------------

    src := []int{10, 20, 30}
    dst := make([]int, len(src))

    count := copy(dst, src)
    fmt.Println(src, dst, count) // → [10 20 30] [10 20 30] 3

    _ = count

    //-------------------------------------------------
    // APPENDING A SLICE (ELLIPSIS ...)
    //-------------------------------------------------
    //
    // `...` expands a slice into individual values.
    //-------------------------------------------------

    numbers := []int{2, 3}
    numbers = append(numbers, 10, 20)
    n := []int{100, 200}
    numbers = append(numbers, n...) // expands slice
    fmt.Println(numbers)
}
