package main

import "fmt"

func Max(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	max := numbers[0]

	for i, v := range numbers {
		if i == 0 {
			continue
		}
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	numbers := []int{2, 3, 5, 7, 7, 8}
	results := Max(numbers)
	fmt.Println(results)
}
