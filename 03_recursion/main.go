package main

import "fmt"

func factorial(n int) int {
	if n == 1 {
		return n
	}
	return n * factorial(n-1)
}

func countdown(n int) int {
	fmt.Println(n)
	if n <= 1 {
		return n
	}
	return countdown(n - 1)
}

func main() {
	fac := factorial(5)
	fmt.Println(fac)

	countdown(19)
}
