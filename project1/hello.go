package main

import (
	"errors"
	"fmt"
	"math"
)

// structure in go
type person struct {
	name string
	age  int
}

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

	// while loop in go
	y := 0
	for y < 5 {
		fmt.Println(y)
		y++
	}

	// iterating over array
	arr1 := []string{"a", "b", "c"}

	for index, value := range arr1 {
		fmt.Println("index", index, value)
	}

	// iterating over map
	ms := make(map[string]string)
	ms["a"] = "anish"
	ms["r"] = "rohan"

	for key, value := range ms {
		fmt.Println("Key", key, "Value : ", value)
	}

	// calling functions
	result := sum(2, 3)
	fmt.Println(result)

	result0, err := sqrt(-16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result0)
	}

	// using struct
	p := person{name: "Anish", age: 23}
	fmt.Println(p)
	fmt.Println(p.name)

	// pointers
	j := 7
	fmt.Println(j, &j)

	// pointers in functions
	k := 7
	inc(&k)
	fmt.Println(k)
}

func sum(x int, y int) int {
	return x + y
}

// functions can return multiple things
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined Negative Number")
	}
	return math.Sqrt(x), nil
}

// take an integer and increment its value
// to increment the value we need to pass pointer to that integer
func inc(x *int) {
	*x++ // dereference the pointer and increment the value
	// notice there is no return statement in the functions, we increment the value of the varibale at the given memory location
}

// go does not have exceptions
