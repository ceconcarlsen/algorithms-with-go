package main

import (
	"fmt"

	"algorithms.com/go/module01"
)

func main() {

	//Num in list
	fmt.Printf(
		"%v\n", module01.NumInList([]int{3, 3, 3, 3}, 4),
	)

	fmt.Printf(
		"%v\n", module01.NumInList([]int{3, 3, 4, 3}, 4),
	)

	//Sum
	fmt.Printf(
		"Sum is: %d\n", module01.Sum([]int{3, 3, 4, 3}),
	)
}
