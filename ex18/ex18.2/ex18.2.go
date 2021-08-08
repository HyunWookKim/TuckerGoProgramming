package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	slice2 := append(slice, 4)

	fmt.Println(slice)
	fmt.Println(slice2)

	slice = append(slice, 3, 4, 5, 6, 7)
	fmt.Println(slice)
}
