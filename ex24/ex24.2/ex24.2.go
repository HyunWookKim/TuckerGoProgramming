package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func SumAtoB(n, a, b int) {
	sum := 0
	for i := 0; i <= b; i++ {
		sum += i
	}

	fmt.Printf("[%d] %d부터 %d까지 합계는 %d입니다.\n", n, a, b, sum)
	wg.Done()
}

func main() {
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go SumAtoB(i+1, 1, 1000000000)
	}
	wg.Wait()
}
