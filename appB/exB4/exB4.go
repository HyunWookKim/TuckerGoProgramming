package main

import "fmt"

func main() {
	var slice []int

	slice, allocCnt := append100times(slice)
	fmt.Println("slice1 - allocCnt:", allocCnt, "cap:", cap(slice))

	var slice2 = make([]int, 0, 1000)
	slice, allocCnt = append100times(slice2)
	fmt.Println("slice2 - allocCnt:", allocCnt, "cap:", cap(slice2))
}

func append100times(slice []int) ([]int, int) {
	var lastCap int = cap(slice)
	var allocCnt int = 0

	for i := 0; i < 1000; i++ {
		slice = append(slice, i)
		if lastCap != cap(slice) {
			fmt.Println(">>>>>>> lastcap:", lastCap)
			allocCnt++
			lastCap = cap(slice)
		}
	}

	return slice, allocCnt
}
