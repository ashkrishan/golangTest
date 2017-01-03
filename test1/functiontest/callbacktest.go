package functiontest

import (
	"fmt"
)

func CallbackTest() {
	var xs []int
	xs = receiverfunction([]int{1, 2, 3, 4, 5, 6}, func(n int) bool {
		return n%2 == 0 //test this expression if true or false in callback
	})
	fmt.Println(xs)
}

func receiverfunction(numbers []int, callback func(int) bool) []int {
	var xs []int
	var val int
	for _, val = range numbers {
		if callback(val) {
			xs = append(xs, val)
		}
	}
	return xs
}
