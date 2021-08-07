package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// rand.Seed(time.Now().UnixNano())

	// n := rand.Intn(100)
	// fmt.Println(n)

	t := time.Now()
	rand.Seed(t.UnixNano())

	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(100))
	}
}
