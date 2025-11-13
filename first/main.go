package main

import "fmt"

func add(a int, b int, c int) int {
	return a + b + c
}

// double calls add to compute x + x
func triple(x int) int {
	return add(x, x, x)
}

func main() {
	x := 4
	fmt.Println("triple of", x, "is", triple(x))
}
