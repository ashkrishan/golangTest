package test

import (
	"fmt"
)

var x = "hell text"

// This function test the closures
func Hellfunc() {
	fmt.Printf("%v %T\n", x, x)
	y := 30
	fmt.Printf("%v", y)
}
