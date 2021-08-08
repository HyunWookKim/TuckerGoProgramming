package main

import "fmt"

func main() {
	const C int = 20
	var b int = C * 20

	C = 20
	fmt.Println(&C)
}
