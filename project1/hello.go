package main

import (
	"fmt"
)

func main() {
	/*
		var x int = 5
		or
		x := 5  // using this format no need to specify varibale type
	*/

	// array in go
	// arrays are of fixed size
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	// slices in go
	// slices are of varibale size
	x := 5
	a := []int{1, 2}
	fmt.Println(a)
	a = append(a, 10)
	fmt.Println(a)

	// if else in go
	if x > 2 {
		fmt.Println("Hello")
	} else if x < 5 {
		fmt.Println("Hello 1")
	} else {
		fmt.Println("NO")
	}

	// maps in go
	vertex := make(map[string]int)
	vertex["triangle"] = 2
	vertex["Square"] = 3

	fmt.Println(vertex["triangle"])
	fmt.Println(vertex)
	delete(vertex, "Square")
	fmt.Println(vertex)

	// for loop in go (only loop in go)
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
