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

After much time spent scratching my head and writing down 0's and 1's on paper, I eventually realized that I could use the properties of binary numbers and bitwise operations to get to a solution.

I wrote down many test cases and even came across some really cool bit twiddling hacks here:
https://graphics.stanford.edu/~seander/bithacks.html

After trying out various techniques and moving bits to the left and the right, I still wasn't able to get a solution that worked 100% of the time. I was close, very close.

Eventually my wandering eyes got the best of me, and I came across a technique to solve this problem online called the Russian Peasant problem. 

https://www.geeksforgeeks.org/russian-peasant-multiply-two-numbers-using-bitwise-operators/

I can't unsee this, so we'll use the Russian Peasant Technique to solve the problem and then have a bake off between this and the naive solution. 

As I've been learning golang, we'll use the native profiling methods to help us figure out which technique is the fastest.

First, let's talk about the Russian Peasant problem and how it works:

Russian Peasant Method Pseudocode from Geeks for Geeks:
```
Let the two given numbers be 'a' and 'b'
1) Initialize result 'res' as 0.
2) Do following while 'b' is greater than 0
   a) If 'b' is odd, add 'a' to 'res'
   b) Double 'a' and halve 'b'
3) Return 'res'. 
```

Let's try using our example of 6 * 7, which is 42!


```
Initial condition: a = 6, b = 7, result = 0

a               b               result (base 10)
0000 0110       0000 0111       0

Iteration 1:
Step 1:
b is > 0 and b is odd, so we will add a to result (6 + 0)

a               b               result (base 10)
0000 0110       0000 0111       (6)

Step 2:
double a which can be accomplished by left shifting a by 1 or mulitplying by 2

a               b               result (base 10)
0000 1100       0000 0111       (6)

Step 3:
half b which can be accomplished by right shifting b by 1 or dividing by 2

a               b               result (base 10)
0000 1100       0000 0011       (6)

Iteration 2:
Step 1:
b is still > 0 and b is also odd, so add a to result (12 + 6)

a               b               result (base 10)
0000 1100       0000 0011       (18)

Step 2:
double a
a               b               result (base 10)
0001 1000       0000 0011       (18)

Step 3:
half b
a               b               result (base 10)
0001 1000       0000 0001       (18)

Iteration 3:
Step 1:
b is still > 0 and b is also odd, so add a to result (24 + 18)
a               b               result (base 10)
0001 1000       0000 0001       (42)

Step 2:
double a
a               b               result (base 10)
0011 1000       0000 0001       (42)

Step 3:
half b
a               b               result (base 10)
0011 1000       0000 0000       (42)

Finally the loop terminates since b is 0.

Now we return the secret of the universe, the great question of life and everything:
42

```

Forty Two!!!


https://www.youtube.com/watch?v=5ZLtcTZP2js


Now We'll update the tests and add test cases for negative numbers as well.
Our test case now passes and all is well:
```
➜  src git:(master) ✗ go test
PASS
ok      _/Users/andre/Documents/Code/golang/go-multiply/src     0.007s
```

Let's find out which solution is faster and if there are any advantages to using the Russian Peasant method over using the naive solution.

To measure how long our code takes to run, we can take advantage of golang's defer keyword.
This article here shows us how we can set up our code for profiling:

https://coderwall.com/p/cp5fya/measuring-execution-time-in-go

After adding in and implementing the profiling code, we can now run our tests to see which method is the fastest.

```
➜  src git:(master) ✗ go test -v
=== RUN   TestMultiplyNaivePositiveAndPositive
2019/07/03 20:56:33 MultiplyNaive took 504ns
--- PASS: TestMultiplyNaivePositiveAndPositive (0.00s)
        multiply_test.go:6: Testing MultiplyNaive 6 * 7
=== RUN   TestMultiplyNaiveNegativeAndNegative
2019/07/03 20:56:33 MultiplyNaive took 214ns
--- PASS: TestMultiplyNaiveNegativeAndNegative (0.00s)
        multiply_test.go:16: Testing MultiplyNaive -6 * -7
=== RUN   TestMultiplyNaiveNegativeAndPositive
2019/07/03 20:56:33 MultiplyNaive took 154ns
--- PASS: TestMultiplyNaiveNegativeAndPositive (0.00s)
        multiply_test.go:25: Testing MultiplyNaive -6 * 7
=== RUN   TestMultiplyNaivePositiveAndNegative
2019/07/03 20:56:33 MultiplyNaive took 135ns
--- PASS: TestMultiplyNaivePositiveAndNegative (0.00s)
        multiply_test.go:34: Testing MultiplyNaive 6 * -7
=== RUN   TestMultiplyBitwisePositiveAndPositive
2019/07/03 20:56:33 MultiplyBitwise took 178ns
--- PASS: TestMultiplyBitwisePositiveAndPositive (0.00s)
        multiply_test.go:44: Testing MultiplyBitwise 6 * 7
=== RUN   TestMultiplyBitwiseNegativeAndNegative
2019/07/03 20:56:33 MultiplyBitwise took 150ns
--- PASS: TestMultiplyBitwiseNegativeAndNegative (0.00s)
        multiply_test.go:53: Testing MultiplyBitwise -6 * -7
=== RUN   TestMultiplyBitwiseNegativeAndPositive
2019/07/03 20:56:33 MultiplyBitwise took 209ns
--- PASS: TestMultiplyBitwiseNegativeAndPositive (0.00s)
        multiply_test.go:62: Testing MultiplyBitwise -6 * 7
=== RUN   TestMultiplyBitwisePositiveAndNegative
2019/07/03 20:56:33 MultiplyBitwise took 129ns
--- PASS: TestMultiplyBitwisePositiveAndNegative (0.00s)
        multiply_test.go:71: Testing MultiplyBitwise 6 * -7
PASS
ok      _/Users/andre/Documents/Code/golang/go-multiply/src     0.007s
```

The results above are from just one test run, but we can see that both methods of solving the problem are neck and neck in terms of execution time.

There is one outlier, which is the 504 nanosecond run of the first function under test. I suspect this is largely due to golang starting up it's testing engine for us to run against.

Which function wins?

Both MultiplyNaive and MultiplyBitwise accomplish the same task. They even run in about the same time. To me, the clear winner is MultiplyNaive. The code is simple and easy to understand. MultiplyBitwise is clever, yet the code is harder to reason about. In the future, if I or one of my team mates were to have to come back and maintain this code, MultiplyNaive would be much easier to change.

Just like with many problems in Computer Science and Software Engineering, there are many ways to solve a problem. Sometimes the simple solutions are the best :)



