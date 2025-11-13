package main

import "fmt"

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	prev1 := 0
	prev2 := 1

	for i := 2; i <= n; i++ {
		current := prev1 + prev2
		prev1 = prev2
		prev2 = current
	}
	return prev2
}

func main() {
	fmt.Println(Fibonacci(8))
}
