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


