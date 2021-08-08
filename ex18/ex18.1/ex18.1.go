package main

import "fmt"

func main() {
	//var slice []int
	slice := []int{1, 2, 3}
	// slice := make([]int, 3)

	if len(slice) == 0 {
		fmt.Println("slice is empty", slice)
	}
	fmt.Println(slice)

	for i := 0; i < len(slice); i++ {
		slice[i] = 10
	}
	fmt.Println(slice)

	for i, v := range slice {
		slice[i] = v * 2
	}
	fmt.Println(slice)
}
