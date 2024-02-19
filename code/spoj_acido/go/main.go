package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var scanner *bufio.Scanner
	if len(os.Args) == 2 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer f.Close()
		scanner = bufio.NewScanner(f)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	for {
		scanner.Scan()
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		letters := strings.Split(line, "")
		count := 0
		stack := make([]string, 0) // empty stack (LIFO)
		for _, letter := range letters {
			if len(stack) == 0 {
				stack = append(stack, letter) // push
			} else {
				topPosition := len(stack) - 1
				topValue := stack[topPosition] // peek, get the top element
				bs := topValue == "B" && letter == "S" || topValue == "S" && letter == "B"
				cf := topValue == "C" && letter == "F" || topValue == "F" && letter == "C"
				if bs || cf {
					stack = stack[:topPosition] // pop, remove top element
					count++
				} else {
					stack = append(stack, letter) // push
				}
			}
		}
		fmt.Println(count)
	}
}
