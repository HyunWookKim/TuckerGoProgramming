package main

import "fmt"

const Pig int = 0
const Cow int = 1
const Chicken int = 2

const (
	C1 uint = iota + 1
	C2
	C3
)

const (
	BitFlag1 uint = 1 << iota
	BitFlag2
	BitFlag3
)

func PrintAnimal(animal int) {
	if animal == Pig {
		fmt.Println("꿀꿀")
	} else if animal == Cow {
		fmt.Println("음메")
	} else if animal == Chicken {
		fmt.Println("꼬끼오")
	} else {
		fmt.Print("...")
	}
}

func main() {
	PrintAnimal(Pig)
	PrintAnimal(Cow)
	PrintAnimal(Chicken)

	fmt.Println("C1", C1, "C2", C2, "C3", C3)
	fmt.Println(BitFlag1, BitFlag2, BitFlag3)
	fmt.Printf("%08b, %08b, %08b", BitFlag1, BitFlag2, BitFlag3)
}
