package sample

import (
	"fmt"
)

var debug = false

func EnableDebug() {
	debug = true
}

func Sum(a, b int) int {
	if debug {
		fmt.Printf("a: %v, b: %v\n", a, b)
	}
	return a + b
}

func DoSum(a, b int) int {
	return a + b
}
