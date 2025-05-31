package module01

func NumInList(list []int, num int) bool {
	var isInList = false

	for i := 0; i < len(list); i++ {
		if list[i] == num {
			isInList = true
			break
		}
	}
	return isInList
}
