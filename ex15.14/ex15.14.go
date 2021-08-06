package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	str1 := "Hello 월드"
	str2 := str1

	// str1 = "Hello"
	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))

	fmt.Println(stringHeader1)
	fmt.Println(stringHeader2)
}
