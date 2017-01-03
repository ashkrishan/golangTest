package functiontest

import (
	"fmt"
)

func changethisstring(z []string) {
	z[0] = "Rbs"
	fmt.Printf("%v %v %p %p\n", z, z[0], &z, &z[0])
}

func StringMakeTest() {
	xs := make([]string, 1, 25)
	fmt.Printf("%v %p\n", xs, &xs)
	changethisstring(xs)
	fmt.Printf("%p\n", &xs)
}
