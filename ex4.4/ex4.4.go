package main

import "fmt"

func main() {

	a := 3 //int - 64bit = int 64
	var b float64 = 3.5

	var c int = int(b)

	d := float64(a) * b //3.5 -> 3 (소수점은 버림)

	var e int64 = 7
	f := a * int(e)

	fmt.Println(a, b, c, d, e, f)
}
