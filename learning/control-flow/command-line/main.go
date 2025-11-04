
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

   

        fmt.Println("Too Expensive ", os.Args)
		fmt.Println("path: ", os.Args[0])
		fmt.Println("first element: ", os.Args[1])
		fmt.Printf("fn: %T\n ", os.Args[2])

		fmt.Println("sec", len(os.Args))
		var r, e = strconv.ParseFloat(os.Args[2], 32)

		fmt.Printf("%T\n", r)
		_= e

  
}
 