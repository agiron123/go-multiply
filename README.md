# Multiply Coding Challenge

### Prompt: 

 Write a function in any language that multiplies any two integers without using the multiplication or division operators. Aim for the most the elegant code you can. We've done this problem ourselves several times, and there's a big difference between your first attempt and how nice it can look after you try to simplify it. That said, be assured that there's no "gotcha" aspect to this question – no obscure trick that you either discover or not – and no particular "right" formulation. We just want to know what you think "elegant" code looks like.


## Solution and Writeup

This is my first time using go, so bear with me as I'm learning as I go (pun intended).

I will be using this section to communicate my approach to solving the problem and showing my process every step of the way.

First, I started out on paper and wrote out the prompt then tried to figure out a naive solution to the problem.

Let's see if we can write out a simple function to figure out the answer to life, the universe and everything.
The Naive approach to this problem can be found in MultiplyNaive in main.go.

I've set up the project to include a failing test so that we can test drive this solution as we go.
Our function initially returns 41, which is clearly off by one.

Running go test, we should get:

```
--- FAIL: TestMultiply (0.00s)
        multiply_test.go:10: The answer to the great question of life, the universe and everything is ... 42
FAIL
exit status 1
FAIL    _/Users/andre/Documents/Code/golang/go-multiply/src     0.006s
```

Using a for loop and an accumulator variable, we can come up with a pretty simple solution:
```
func MultiplyNaive(a, b int)(product int) {
	var result int = 0
	for i := 0; i < b; i++ {
		result += a
	}
	return result
}
```

It doesn't work for negative numbers just yet, so we'll add some additional test cases to test:
```
-6 * 7
6 * -7
-6 * -7
```

Now that the test cases are added, we can now implement our naive solution:
```
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
```

Our naive solution is now working and our tests are passing!

```
➜  src git:(2-fix-multiplynaive-failing-test) ✗ go test -v
=== RUN   TestMultiply
--- PASS: TestMultiply (0.00s)
        multiply_test.go:6: Testing Multiply 6 * 7
=== RUN   TestMultiplyNegativeAndNegative
--- PASS: TestMultiplyNegativeAndNegative (0.00s)
        multiply_test.go:16: Testing Multiply -6 * -7
=== RUN   TestMultiplyNegativeAndPositive
--- PASS: TestMultiplyNegativeAndPositive (0.00s)
        multiply_test.go:26: Testing Multiply -6 * 7
=== RUN   TestMultiplyPositiveAndNegative
--- PASS: TestMultiplyPositiveAndNegative (0.00s)
        multiply_test.go:36: Testing Multiply 6 * -7
PASS
ok      _/Users/andre/Documents/Code/golang/go-multiply/src     0.006s
```


Alright, so the naive solution is implemented. Can we do better?

The Computer Scientist in me thinks we can try to use Bit twiddling and figure out a solution.

Let's first create a new function, MultiplyBitwise and add a failing test.
```
func MultiplyBitwise(a,b int) (product int) {
		return 41
}
```

```
func TestMultiplyBitwise(t *testing.T) {
		result := MultiplyBitwise(6,7)

		if result != 42 {
			t.Errorf("Expected 42")
		}
}
```

Now that our test is in order, let's see if we can implement by taking advantage of Bitwise operators and the properties of binary numbers.