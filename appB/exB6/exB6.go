package main

import (
	"fmt"
	"sort"
)

type IntSlice []int

func (p IntSlice) Len() int {
	return len(p)
}

func (p IntSlice) Less(i, j int) bool {
	return p[i] < p[j] //오름 차순 변환
}

func (p IntSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	s := []int{5, 2, 3, 6, 1, 7, 9, 11}
	// sort.Sort(sort.IntSlice(s))
	sort.Ints(s)
	fmt.Println(s)
}
