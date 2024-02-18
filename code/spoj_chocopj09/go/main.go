package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
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
