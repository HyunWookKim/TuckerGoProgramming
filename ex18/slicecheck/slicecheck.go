package main

import "fmt"

// func addNum(slice *[]int) {
// 	*slice = append(*slice, 4)
// }

func addNum(slice []int) []int {
	slice = append(slice, 4)
	return slice
}

func main() {
	slice := []int{1, 2, 3}
	// addNum(&slice)
	slice = addNum(slice)
	fmt.Println(slice)
}
