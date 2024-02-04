// This section will dissect a Dynamic Programming problem called Fibonacci sequence to illustrate the above 2 very important concepts.
//
// Fibonacci sequence is a well-known sequence of numbers that looks something like this 0,1,1,2,3,5,8,13, 21,…..
//
// fib(n) = fib(n-1) + fib(n-2) where fib(0) = 0 and fib(1) = 1
package fibonnaci

import (
	"fmt"
	"time"
)

// Time complexity: O(n^2)
func withRecursion(value int) int {
	if value <= 2 {
		return 1
	}
	prev := withRecursion(value - 1)      //O(n)
	prev_prev := withRecursion(value - 2) //O(n)
	return prev_prev + prev
}

// Time complexity: O(n)
func withRecursionAndMemoization(value int, memo map[int]int) int {
	if memo[value] > 0 {
		return memo[value]
	}
	if value <= 2 {
		return 1
	}
	prev := withRecursionAndMemoization(value-1, memo)      // O(n)
	prev_prev := withRecursionAndMemoization(value-2, memo) // O(n-1)
	memo[value] = prev_prev + prev
	return memo[value]
}

// Time Complexity O(n-2) => O(n)
func withTabulation(value int) int {
	numberOfFibonacciItems := value + 1
	array := make([]int, numberOfFibonacciItems)

	array[0] = 0
	array[1] = 1
	for index := 2; index < numberOfFibonacciItems; index++ { //O(n)
		prev := array[index-1]
		prev_prev := array[index-2]
		array[index] = prev_prev + prev
	}
	return array[value]
}

func DoFibonacci(value int) {
	var now time.Time
	var result int
	inputs := [...]int{2, 7, 13, 37}
	for _, value := range inputs {
		now = time.Now()
		result = withRecursion(value)
		fmt.Printf("B: input %d result %d in %d time \n", value, result, time.Since(now).Microseconds())

		now = time.Now()
		result = withRecursionAndMemoization(value, make(map[int]int))
		fmt.Printf("M: input %d result %d in %d time \n", value, result, time.Since(now).Microseconds())

		now = time.Now()
		result = withTabulation(value)
		fmt.Printf("T: input %d result %d in %d time \n", value, result, time.Since(now).Microseconds())

		println()
	}
}
