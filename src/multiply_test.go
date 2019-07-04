package main

import "testing"

func TestMultiplyNaivePositiveAndPositive(t *testing.T) {
	t.Log("Testing MultiplyNaive 6 * 7")
	result := MultiplyNaive(6, 7)

	// Obligatory: https://www.youtube.com/watch?v=5ZLtcTZP2js
	if result != 42 {
		t.Errorf("The answer to the great question of life, the universe and everything is ... 42")
	}
}

func TestMultiplyNaiveNegativeAndNegative(t *testing.T) {
	t.Log("Testing MultiplyNaive -6 * -7")
	result := MultiplyNaive(-6, -7)

	if result != 42 {
		t.Errorf("The answer to the great question of life, the universe and everything is ... 42")
	}
}

func TestMultiplyNaiveNegativeAndPositive(t *testing.T) {
	t.Log("Testing MultiplyNaive -6 * 7")
	result := MultiplyNaive(-6, 7)

	if result != -42 {
		t.Errorf("Expected -42")
	}
}

func TestMultiplyNaivePositiveAndNegative(t *testing.T) {
	t.Log("Testing MultiplyNaive 6 * -7")
	result := MultiplyNaive(6, -7)

	if result != -42 {
		t.Errorf("Expected -42")
	}
}

// Multiply Bitwise tests
func TestMultiplyBitwisePositiveAndPositive(t *testing.T) {
		t.Log("Testing MultiplyBitwise 6 * 7")
		result := MultiplyBitwise(6,7)

		if result != 42 {
			t.Errorf("Expected 42")
		}
}

func TestMultiplyBitwiseNegativeAndNegative(t *testing.T) {
	t.Log("Testing MultiplyBitwise -6 * -7")
	result := MultiplyBitwise(-6, -7)

	if result != 42 {
		t.Errorf("The answer to the great question of life, the universe and everything is ... 42")
	}
}

func TestMultiplyBitwiseNegativeAndPositive(t *testing.T) {
	t.Log("Testing MultiplyBitwise -6 * 7")
	result := MultiplyBitwise(-6, 7)

	if result != -42 {
		t.Errorf("Expected -42")
	}
}

func TestMultiplyBitwisePositiveAndNegative(t *testing.T) {
	t.Log("Testing MultiplyBitwise 6 * -7")
	result := MultiplyBitwise(6, -7)

	if result != -42 {
		t.Errorf("Expected -42")
	}
}