package functiontest

import (
	"fmt"
)

func receivepointer(xPtr *int) {
	fmt.Println("xPtr ", xPtr, *xPtr, &xPtr)
	*xPtr = 0
	fmt.Println("xPtr ", xPtr, *xPtr, &xPtr)
}

func PointerTest() {
	var x int
	x = 12
	fmt.Println(x, &x)
	receivepointer(&x)
	fmt.Println(x, &x)
}
