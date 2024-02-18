package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	scanner.Scan()

	input := strings.Split(scanner.Text(), " ")
	numberOfBalls, _ := strconv.Atoi(input[0])
	maximumBallsPerTime, _ := strconv.Atoi(input[1])

	if maximumBallsPerTime >= numberOfBalls {
		fmt.Print("Paula")
	}

	value := numberOfBalls%(maximumBallsPerTime+1) != 0
	if value {
		fmt.Print("Paula")
	} else {
		fmt.Print("Carlos")
	}

}
