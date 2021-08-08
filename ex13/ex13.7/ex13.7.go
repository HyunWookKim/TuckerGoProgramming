package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	A int8 //1byte
	B int  //8byte
	C int8 //1byte
	D int  //8byte
	E int8 //1byte
}

type User2 struct {
	A int8 //1byte
	C int8 //1byte
	E int8 //1byte
	B int  //8byte
	D int  //8byte

}

func main() {
	user := User{1, 2, 3, 4, 5}
	user2 := User2{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user))
	fmt.Println(unsafe.Sizeof(user2))
}
