//-------------------------------------------------
// Constants and IOTA
//-------------------------------------------------
//
// This example demonstrates Go constants — typed, untyped,
// grouped constants, and the use of the special constant generator `iota`.
//
// Ethereum-inspired variable names are used for contextual clarity.
//-------------------------------------------------

package main

import "fmt"

func main() {

    //-------------------------------------------------
    // DECLARING CONSTANTS
    //-------------------------------------------------
    // Constants are declared using `const` and must be initialized.
    const maxBlockSize int = 7   // typed constant
    const etherPrice = 3500.75   // untyped constant

    // Go supports only:
    // boolean, rune, integer, floating-point, complex, and string constants.

    //-------------------------------------------------
    // GROUPED CONSTANTS
    //-------------------------------------------------
    const (
        maxGasLimit       = 30000000  // untyped constant
        blockReward float64 = 2.0     // typed constant
    )

    const (
        shardCount     = 4
        validatorCount = 5
    )
	_,_,_,_=shardCount, maxGasLimit,maxBlockSize,etherPrice
    // Grouped constants repeat the previous type and value
    const (
        minDifficulty = 100000
        maxDifficulty // => 100000 (inherits previous value)
        avgDifficulty // => 100000
    )
	_=minDifficulty
    //-------------------------------------------------
    // CONSTANT RULES
    //-------------------------------------------------
    // 1. Constants cannot be reassigned.
    const networkID int = 1
    // networkID = 5 // compile-time error
	_=networkID
    // 2. Constants are compile-time only — not runtime.
    // const gasFee = math.Pow(2, 3) // invalid: runtime function call

    // 3. Cannot initialize a constant with a variable (runtime value).
    nodeCount := 5
    // const activeNodes = nodeCount // not allowed

    // 4. len() is allowed only if its argument is a constant literal.
    const chainNameLength = len("Ethereum") // OK
    chainName := "Ethereum"
    // const invalidLen = len(chainName) // invalid: str is a variable

    _, _ ,_= nodeCount, chainName,chainNameLength

    //-------------------------------------------------
    // UNTYPED CONSTANTS
    //-------------------------------------------------
    const blockCount = 5
    const blockRewardRate float64 = 2.5

    var txCount int = 12
    var gasPrice float64 = 1.1

    fmt.Println(blockCount * blockRewardRate)
    // No error: blockCount is untyped → becomes float64 at use (12.5)

    // fmt.Println(txCount * gasPrice)
    // Error: mismatched types int and float64
    _, _ = txCount, gasPrice

    //-------------------------------------------------
    // IOTA
    //-------------------------------------------------
    // iota is a special constant generator starting at 0
    // and incrementing automatically within a const block.

    const (
        blockA = iota
        blockB
        blockC
    )
    fmt.Println(blockA, blockB, blockC) // => 0 1 2

    //-------------------------------------------------
    //   			IOTA 
    //-------------------------------------------------
    const (
        NetworkMainnet = iota // 0
        NetworkGoerli         // 1
        NetworkSepolia        // 2
        NetworkPrivate        // 3
    )
    _ = NetworkMainnet

    //-------------------------------------------------
    // IOTA WITH STEP (BLOCK INTERVALS)
    //-------------------------------------------------
    const (
        blockInterval10s = iota * 10 // 0
        blockInterval20s             // 10
        blockInterval30s             // 20
    )
    fmt.Println("The value of the blockInterval30s  =",blockInterval30s)
}
