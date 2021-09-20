package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const input = "Now is winter of our discontnent,\nMade glorious summer by this sun of York.\n"
	scanner := bufio.NewScanner(strings.NewReader(input)) //scanner 생성
	scanner.Split(bufio.ScanWords)                        //글자 단위로 검색(스캔)

	count := 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		count++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	fmt.Printf("%d\n", count)
}
