package main

import "fmt"

type OpFn func(int, int) int

// func getOperator(op string) OpFn {
func getOperator(op string) func(int, int) int {
	if op == "+" {
		return func(a, b int) int {
			return a + b
		}
	} else if op == "*" {
		return func(a, b int) int {
			return a * b
		}
	} else {
		return nil
	}
}

func main() {
	// var operator OpFn
	var operator func(int, int) int
	operator = getOperator("*")

	var result = operator(3, 4)
	fmt.Println(result)
}
