package main

import "testing"

func TestMultiply(t *testing.T) {
		result := MultiplyNaive(6, 7)

		// Obligatory: https://www.youtube.com/watch?v=5ZLtcTZP2js
		if result != 42 {
				t.Errorf("The answer to the great question of life, the universe and everything is ... 42")
		}
}