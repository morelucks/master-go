package main

import "fmt"


func main (){
	fmt.Println("Backing array started")

	numbers := []int{10, 15,20, 30, 40}
another := numbers[1:3]
fmt.Println(len(another))
fmt.Println(numbers) 

fmt.Println(cap(another)) 
fmt.Println(another)

}