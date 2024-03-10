package main

import (
	"fmt"
)

func main() {

	//input := [6]int{1, 1, 1, 3, 2, 3}
	input := [3]int{4, 5, 6}

	count := 0
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] == input[j] {
				count++
			}
		}
	}

	fmt.Println(count)
}
