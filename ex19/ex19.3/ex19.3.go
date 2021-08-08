package main

import "fmt"

type account struct {
	balacne   int
	firstname string
	lastname  string
}

func (a1 *account) withdrawPointer(amount int) {
	a1.balacne -= amount
}

func (a2 account) withdrawValue(amount int) {
	a2.balacne -= amount
}

func (a3 account) withdrawReturnValue(amount int) account {
	a3.balacne -= amount
	return a3
}

func main() {
	var mainA *account = &account{100, "joe", "park"}
	mainA.withdrawPointer(30)
	fmt.Println(mainA.balacne)

	mainA.withdrawValue(20)
	fmt.Println(mainA.balacne)

	var mainB account = mainA.withdrawReturnValue(20)
	fmt.Println(mainB.balacne)

	mainB.withdrawPointer(30)
	fmt.Println(mainB.balacne)

	fmt.Println(mainA.balacne)
}
