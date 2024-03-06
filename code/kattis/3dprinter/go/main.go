package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
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

	powInt := func(x, y int) int {
		return int(math.Pow(float64(x), float64(y)))
	}

	for {
		scanner.Scan()
		input := scanner.Text()
		if input == "" {
			break
		}

		statues, _ := strconv.Atoi(input)
		days := 1
		for {
			value := powInt(2, days-1)
			if value >= statues {
				break
			}
			days += 1
		}

		fmt.Println(days)
	}
}
