package functiontest

import (
	"fmt"
)

func RecursionTest() {
	fmt.Println(factorial(4)) // calculting factorial of 4

}

//Recursion must not be used in normal circunstances as it may cause perfomance issues on stack.
//This example is just to understand recursion concept in GO.
func factorial(x int) int {
	if x == 0 {
		return 1 //will exit function with return value 1 upon x reaches 0
	}
	return x * factorial(x-1) //Recurse for instance if value is 4. function will recurse 4,3,2,1 and total value returned will be 24
}
