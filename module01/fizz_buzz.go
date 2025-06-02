package module01

import (
	"fmt"
)

func isFizz(i int) bool {
	return i%3 == 0
}

func isBuzz(i int) bool {
	return i%5 == 0
}

func FizzBuzz(n int) {

	for i := 1; i <= n; i++ {
		if isFizz(i) && isBuzz(i) {
			fmt.Printf("Fizz Buzz\n")
		} else if isFizz(i) {
			fmt.Printf("Fizz\n")
		} else if isBuzz(i) {
			fmt.Printf("Buzz\n")
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}
