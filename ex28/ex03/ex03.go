package main

import (
	"errors"
	"fmt"
	"strings"
)

func Atoi(input string) (int, error) {
	rst := 0
	negative := false
	input = strings.TrimSpace(input)

	if input[0] == '-' {
		negative = true
		input = input[1:]
	}

	for _, c := range input {
		if c >= '0' && c <= '9' {
			rst *= 10
			rst += int(c - '0')
		} else {
			err := fmt.Sprintf("cannot convert to int, this string has a '%c'", c)
			return 0, errors.New(err)
		}
	}

	if negative {
		rst *= -1
	}

	return rst, nil
}
