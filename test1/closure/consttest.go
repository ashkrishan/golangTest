package closure

import (
	"fmt"
)

const (
	_  = iota
	KB = 1 << (iota * 10) //1 << (1 * 10)
	MB = 1 << (iota * 10) //1 << (2 * 10)
)

func Testiota() {
	fmt.Printf("%b\t-%d\n", KB, KB)
	fmt.Printf("%b\t-%d\n", MB, MB)
	x := 32
	y := x + 20
	fmt.Printf("%v is type %T", x, x)
	fmt.Printf("%v is type %T", y, y)
}
