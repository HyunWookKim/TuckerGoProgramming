package main

import (
	"fmt"
)

func main() {
	var array [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var slice1 []int = array[1:5]
	var slice2 []int = slice1[1:8:9]
	var slice3 []int = make([]int, 5)
	var slice4 []int = make([]int, 0)
	var slice5 []int = []int{1, 2, 3, 4, 5}
	var slice6 []int

	// array[1] = 11
	// array[8] = 11
	// slice1[3] = 11
	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)

	fmt.Println("slice3", slice3)
	fmt.Println("slice4", slice4)
	if slice4 != nil {
		fmt.Println("slice4 is not nil")
	}
	fmt.Println("slcie5", slice5)
	fmt.Println("slcie6", slice6)
	if slice6 != nil {
		fmt.Println("slice6 is not nil")
	}
}
