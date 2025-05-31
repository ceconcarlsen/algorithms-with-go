package module01

func Sum(numbers []int) int {
	var sum = 0

	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	return sum
}
