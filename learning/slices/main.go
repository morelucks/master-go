package main

import "fmt"


//-------------------------------------------------
// Go Slices Explained
//-------------------------------------------------
//
// This example explores how Go handles *slices* — 
// one of the most powerful and flexible data types 
// in Go. Slices build upon arrays and are heavily 
// used in blockchain development for managing dynamic 
// datasets like transaction lists, wallet balances, 
// and validator sets.
//
// A slice is a lightweight data structure describing 
// a contiguous section of an underlying array.
//
// Slice Header Fields:
//  - Pointer : address of the first element
//  - Length  : number of accessible elements
//  - Capacity: max elements until reallocation
//-------------------------------------------------

func main (){
    //-------------------------------------------------
    // DECLARATION AND INITIALIZATION
    //-------------------------------------------------
    // A nil slice has no underlying array.
    var txHashes []string
    fmt.Printf("txHashes == nil: %v, value: %#v\n", txHashes == nil, txHashes)
    // Output: true, []string(nil)

    // You can create slices using literals.
    recentBlocks := []int{18234001, 18234002, 18234003, 18234004}
    fmt.Printf("recentBlocks: %v (len=%d, cap=%d)\n", recentBlocks, len(recentBlocks), cap(recentBlocks))

    // Or create using make() with defined length (and optional capacity).
    gasPool := make([]uint64, 3)
    fmt.Printf("gasPool initialized: %v\n", gasPool) // => [0 0 0]

    //-------------------------------------------------
    // SLICING ARRAYS
    //-------------------------------------------------
    // Arrays, slices, and strings are *sliceable*.
    // Slicing doesn’t copy data—it creates a view.
    allBalances := [5]float64{5.0, 8.5, 2.3, 1.2, 10.0}

    activeBalances := allBalances[1:4] // start inclusive, end exclusive
    fmt.Printf("activeBalances: %v\n", activeBalances) // => [8.5 2.3 1.2]

    // Omitting start or end index uses defaults.
    fmt.Println(allBalances[:3]) // [5 8.5 2.3]
    fmt.Println(allBalances[2:]) // [2.3 1.2 10]
    fmt.Println(allBalances[:])  // full slice [5 8.5 2.3 1.2 10]

    //-------------------------------------------------
    // APPENDING ELEMENTS
    //-------------------------------------------------
    // Append dynamically increases slice length.
    mempool := []string{"tx01", "tx02"}
    mempool = append(mempool, "tx03", "tx04")
    fmt.Println("Mempool transactions:", mempool)

    // Append can also merge slices.
    pendingTx := []string{"tx05", "tx06"}
    mempool = append(mempool, pendingTx...)
    fmt.Println("Extended mempool:", mempool)

    //-------------------------------------------------
    // SHARED BACKING ARRAY
    //-------------------------------------------------
    // Two slices from the same array share memory!
    validators := []string{"Lido", "Binance", "Coinbase", "RocketPool"}
    v1 := validators[:2]
    v2 := validators[1:3]
    fmt.Println("Before modification:", validators, v1, v2)

    v1[1] = "StakeWise" // modifies the backing array
    fmt.Println("After modification:", validators, v1, v2)

    //-------------------------------------------------
    // COPYING SLICES SAFELY
    //-------------------------------------------------
    // Use copy() to duplicate data into a new slice.
    original := []int{10, 20, 30}
    backup := make([]int, len(original))
    copy(backup, original)
    fmt.Println("original:", original, "backup:", backup)

    original[0] = 99
    fmt.Println("After modifying original:", original, "backup remains:", backup)

    //-------------------------------------------------
    // COMPARING SLICES
    //-------------------------------------------------
    // Slices cannot be compared directly (except to nil).
    var a, b = []int{1, 2, 3}, []int{1, 2, 3}
    equal := len(a) == len(b)
    for i := range a {
        if a[i] != b[i] {
            equal = false
            break
        }
    }
    fmt.Printf("Are slices equal? %v\n", equal)

    //-------------------------------------------------
    // SLICE GROWTH AND CAPACITY
    //-------------------------------------------------
    fmt.Println("\nSlice growth demonstration:")
    txQueue := []string{}
    for i := 1; i <= 5; i++ {
        txQueue = append(txQueue, fmt.Sprintf("tx#%d", i))
        fmt.Printf("len=%d, cap=%d, data=%v\n", len(txQueue), cap(txQueue), txQueue)
    }



	// code:=[]byte{}
	// code=append(code, 43)
	// fmt.Println(code)
	// remaincode:=[]byte{20,54,65,50,4}

	// code = append(code, remaincode...)
	// newlocation:=[]byte{}
	// copy(newlocation, code)
	// fmt.Println(newlocation)

	// f3code, s2lcode:=code[:3], code[2:]

	// fmt.Println(f3code)
	// fmt.Println(s2lcode)

	// f3code[2]=90

	// fmt.Println(f3code)
	// fmt.Println(s2lcode)//?????

}