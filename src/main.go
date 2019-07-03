// Andre Giron 2019

// Package multiply multiplies two integers without using the multiplication or division operators.
package main

import "fmt"

func MultiplyNaive(a, b int) (product int) {
	// First, check the signs of a and b.
	isPositive := true
	if (a < 0 && b > 0) || (a > 0 && b < 0) { // Result is going to be negative
		isPositive = false
	}

	// Set both b and a to be positive integers.
	// Cheating for now, but we'll discuss how to do this with bitwise operators very soon!
	if a < 0 {
		a = -a
	}

	if b < 0 {
		b = -b
	}

	var result int = 0
	for i := 0; i < b; i++ {
		result += a
	}

	if !isPositive {
		return -result
	}

	return result
}

func main() {
	fmt.Println("6 * 7 is:", MultiplyNaive(6, 7))
}
