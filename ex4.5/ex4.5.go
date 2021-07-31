package main

import "fmt"

func main() {
	var a int16 = 3456   //a는 int 16타입 - 2바이트 정수
	var b int8 = int8(a) //int16 -> int 8

	fmt.Println(a, b)
}
