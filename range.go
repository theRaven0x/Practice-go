package piscine

func CountGreaterThanFive() int {

	values := []int{3, 9, 2, 0, 11, 4}
	count := 0

	for _, v := range values {
		if v > 5 {
			count = count + 1
		}
	}
	return count
}
