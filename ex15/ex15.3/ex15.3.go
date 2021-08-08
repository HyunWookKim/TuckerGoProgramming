package main

import "fmt"

func main() {
	str := "Hello 월드"
	arr := []rune(str)
	//rune => int32 별칭 타입
	//rune은 한칸이 4바이트

	for i := 0; i < len(arr); i++ {
		fmt.Printf("타입 : %T, 값 : %d, 문자값 : %c\n", arr[i], arr[i], arr[i])
	}
}
