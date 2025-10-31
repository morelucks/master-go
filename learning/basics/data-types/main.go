//-------------------------------------------------
// Go Data Types 
//-------------------------------------------------
//
// This example explores Go's core data types using
// Ethereum-related concepts like blocks, gas prices,
// balances, miners, and wallets.
//
// Each type determines a set of possible values and
// operations specific to those values.
//-------------------------------------------------

package main

import "fmt"

func main() {

    //-------------------------------------------------
    // NUMERIC TYPES
    //-------------------------------------------------
    // Integer and floating-point types:
    //
    // Unsigned integers:
    //  uint8   : 0 to 255
    //  uint16  : 0 to 65535
    //  uint32  : 0 to 4294967295
    //  uint64  : 0 to 18446744073709551615
    //
    // Signed integers:
    //  int8    : -128 to 127
    //  int16   : -32768 to 32767
    //  int32   : -2147483648 to 2147483647
    //  int64   : -9223372036854775808 to 9223372036854775807
    //
    // Floating-point:
    //  float32 : IEEE-754 32-bit floating-point numbers
    //  float64 : IEEE-754 64-bit floating-point numbers
    //
    // Complex numbers:
    //  complex64  : float32 real and imaginary parts
    //  complex128 : float64 real and imaginary parts
    //
    // Aliases:
    //  byte : alias for uint8
    //  rune : alias for int32

    var blockHeight int64 = 21_000_000 // example of signed integer
    fmt.Printf("blockHeight type: %T, value: %v\n", blockHeight, blockHeight)

    var gasPrice uint32 = 120_000_000 // unsigned integer for gas in gwei
    fmt.Printf("gasPrice type: %T, value: %v\n", gasPrice, gasPrice)

    var difficulty float64 = 1.2345e12 // floating-point for large network values
    fmt.Printf("difficulty type: %T, value: %v\n", difficulty, difficulty)

    //-------------------------------------------------
    // FLOAT TYPES
    //-------------------------------------------------
    var ethPriceUSD, marketChange, stakingReward float64 = 3320.45, -2.5, 5.0
    fmt.Printf("ethPriceUSD: %T, marketChange: %T, stakingReward: %T\n",
        ethPriceUSD, marketChange, stakingReward)

    //-------------------------------------------------
    // RUNE TYPE
    //-------------------------------------------------
    var network rune = 'M' // 'M' for Mainnet
    fmt.Printf("network type: %T, value (decimal): %v\n", network, network)
    fmt.Printf("network value (hex): %x\n", network)
    fmt.Printf("network as character: %c\n", network)

    //-------------------------------------------------
    // BOOLEAN TYPE
    //-------------------------------------------------
    var minerActive bool = true
    fmt.Printf("minerActive type: %T, value: %v\n", minerActive, minerActive)

    //-------------------------------------------------
    // STRING TYPE
    //-------------------------------------------------
    var walletAddress string = "0xAbC123...9fE"
    fmt.Printf("walletAddress type: %T, value: %q\n", walletAddress, walletAddress)

    //-------------------------------------------------
    // ARRAY TYPE
    //-------------------------------------------------
    // Arrays have a fixed size.
    var recentBlocks = [4]int{18234001, 18234002, 18234003, 18234004}
    fmt.Printf("recentBlocks type: %T, value: %v\n", recentBlocks, recentBlocks)

    //-------------------------------------------------
    // SLICE TYPE
    //-------------------------------------------------
    // Slices are dynamic and more flexible than arrays.
    var topValidators = []string{"Lido", "Coinbase", "Binance", "RocketPool"}
    fmt.Printf("topValidators type: %T, value: %v\n", topValidators, topValidators)

    //-------------------------------------------------
    // MAP TYPE
    //-------------------------------------------------
    // Maps represent key-value pairs.
    walletBalances := map[string]float64{
        "0xAlice": 5.35,
        "0xBob":   12.90,
        "0xCarol": 0.89,
    }
    fmt.Printf("walletBalances type: %T, value: %v\n", walletBalances, walletBalances)

    //-------------------------------------------------
    // STRUCT TYPE
    //-------------------------------------------------
    // Structs group related data under one name.
    type Block struct {
        number   int
        miner    string
        gasUsed  uint64
        reward   float64
    }

    var latestBlock = Block{
        number: 18234001,
        miner:  "0xMinerNode01",
        gasUsed: 29_450_000,
        reward: 2.05,
    }

    fmt.Printf("latestBlock type: %T, value: %+v\n", latestBlock, latestBlock)

    //-------------------------------------------------
    // POINTER TYPE
    //-------------------------------------------------
    var totalTx int = 15_000_000
    ptr := &totalTx
    fmt.Printf("ptr type: %T, value: %v (points to totalTx)\n", ptr, ptr)

    //-------------------------------------------------
    // FUNCTION TYPE
    //-------------------------------------------------
    fmt.Printf("syncNode type: %T\n", syncNode)
}

// simulates a node synchronization process
func syncNode() {
    fmt.Println("Node is synchronizing with the Ethereum network...")
}
