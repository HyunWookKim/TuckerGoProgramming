package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare1(t *testing.T) {

	assert.Equal(t, 82, square(9), "square(9) should be 81")
	// rst := square(9)

	// if rst != 81 {
	// 	t.Errorf("square(9) should be 81 but return %d", rst)
	// }
}

func TestSquare2(t *testing.T) {
	assert.Equal(t, 9, square(3), "square(3) should be 9")
	// rst := square(3)

	// if rst != 9 {
	// 	t.Errorf("square(3) should be 9 but return %d", rst)
	// }
}
