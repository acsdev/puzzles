// Problem Statement
// ++++++++++++++++++
// Given a target sum and an unsorted array of positive integers.
// Whether it is possible to generate the targetSum using the numbers in the array? — canSum(targetSum, integers)
// canSum(8, [2,3,5]) => true, [2,2,2,2] or [2, 3, 3] or [3, 5]
// canSum(8, [3,6,9]) => false
package main

import (
	"fmt"
	"time"
)

func withRecursion(target int, values []int) bool {
	if target == 0 {
		return true
	}
	if target < 0 {
		return false
	}

	for _, value := range values {
		newTarget := target - value
		if withRecursion(newTarget, values) {
			return true
		}
	}
	return false
}

func withRecursionAndMemoization(target int, values []int, memo map[int]bool) bool {
	if memo[target] {
		return memo[target]
	}
	if target == 0 {
		return true
	}
	if target < 0 {
		return false
	}

	for _, value := range values {
		newTarget := target - value
		if withRecursionAndMemoization(newTarget, values, memo) {
			memo[target] = true
			return true
		}
	}
	return false
}

func withTabulation(target int, values []int) bool {
	numberOfItems := target + 1
	canSum := make([]bool, numberOfItems)

	canSum[0] = true

first:
	for index, can := range canSum {
		if can {
			for _, value := range values {
				jump := index + value
				if len(canSum) > jump {
					canSum[jump] = true
					if jump == target {
						fmt.Printf("Jumped in %d", index)
						break first
					}
				}
			}
		}
	}

	return canSum[target]
}

func main() {
	var now time.Time
	var result bool

	inputs := [...]*Pair{
		{target: 8, values: []int{2, 3, 5}},
		{target: 8, values: []int{3, 6, 9}},
	}

	for _, input := range inputs {
		now = time.Now()
		result = withRecursion(input.target, input.values)
		fmt.Printf("B: input %v result %t in %d time \n", input, result, time.Since(now).Nanoseconds())

		now = time.Now()
		result = withRecursionAndMemoization(input.target, input.values, make(map[int]bool))
		fmt.Printf("M: input %v result %t in %d time \n", input, result, time.Since(now).Nanoseconds())

		now = time.Now()
		result = withTabulation(input.target, input.values)
		fmt.Printf("T: input %v result %t in %d time \n", input, result, time.Since(now).Nanoseconds())

		println()
	}
}

type Pair struct {
	target int
	values []int
}
