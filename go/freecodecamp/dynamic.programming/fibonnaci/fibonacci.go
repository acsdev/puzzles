package fibonnaci

import (
	"fmt"
	"time"
)

func withRecursion(value int) int {
	if value <= 2 {
		return 1
	}
	return withRecursion(value-1) + withRecursion(value-2)
}

func withRecursionAndMemoization(value int, memo map[int]int) int {
	if memo[value] > 0 {
		return memo[value]
	}
	if value <= 2 {
		return 1
	}
	memo[value] = withRecursionAndMemoization(value-1, memo) + withRecursionAndMemoization(value-2, memo)
	return memo[value]
}

func withTabulation(value int) int {
	numberOfFibonacciItems := value + 1
	array := make([]int, numberOfFibonacciItems)

	array[1] = 1
	for index := 0; index < numberOfFibonacciItems; index++ {
		currentValue := array[index]
		if (index + 1) < numberOfFibonacciItems {
			array[index+1] += currentValue
		}
		if (index + 2) < numberOfFibonacciItems {
			array[index+2] += currentValue
		}
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
