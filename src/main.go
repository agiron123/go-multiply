// Andre Giron 2019

// Package multiply multiplies two integers without using the multiplication or division operators.
package main

import "fmt"
import "log"
import "time"

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %dns", name, elapsed.Nanoseconds())
}

func MultiplyBitwise(a, b int) (product int) {
		defer timeTrack(time.Now(), "MultiplyBitwise")
		result := 0
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

		// Go does not have while loops! As someone who primarily writes JavaScript and Python, this is neat!
		for b > 0 {
				// Check if b is an odd number
				if b & 1 == 1 {
						result += a
				}

				// Double a. We can either bitshift to the left by 1 or multiply by 2.
				a = a << 1

				// Half b. Just like above, we can either bitshift to the right by 1 or divide by 2.
				b = b >> 1
		}

		if !isPositive {
				return -result
		}

		return result
}

func MultiplyNaive(a, b int) (product int) {
	defer timeTrack(time.Now(), "MultiplyNaive")
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
	fmt.Println("Naive 6 * 7 is:", MultiplyNaive(6, 7))
	fmt.Println("Bitwise 6 * 7 is:", MultiplyBitwise(6,7))
}
