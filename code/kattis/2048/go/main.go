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

	for {
		scanner.Scan()
		possibleLine1 := scanner.Text()
		if possibleLine1 == "" {
			break
		}

		line1 := strings.Split(possibleLine1, " ")

		scanner.Scan()
		line2 := strings.Split(scanner.Text(), " ")

		scanner.Scan()
		line3 := strings.Split(scanner.Text(), " ")

		scanner.Scan()
		line4 := strings.Split(scanner.Text(), " ")

		scanner.Scan()
		direction, _ := strconv.Atoi(scanner.Text())

		var board = [4][4]int{}
		board[0][0], _ = strconv.Atoi(line1[0])
		board[0][1], _ = strconv.Atoi(line1[1])
		board[0][2], _ = strconv.Atoi(line1[2])
		board[0][3], _ = strconv.Atoi(line1[3])

		board[1][0], _ = strconv.Atoi(line2[0])
		board[1][1], _ = strconv.Atoi(line2[1])
		board[1][2], _ = strconv.Atoi(line2[2])
		board[1][3], _ = strconv.Atoi(line2[3])

		board[2][0], _ = strconv.Atoi(line3[0])
		board[2][1], _ = strconv.Atoi(line3[1])
		board[2][2], _ = strconv.Atoi(line3[2])
		board[2][3], _ = strconv.Atoi(line3[3])

		board[3][0], _ = strconv.Atoi(line4[0])
		board[3][1], _ = strconv.Atoi(line4[1])
		board[3][2], _ = strconv.Atoi(line4[2])
		board[3][3], _ = strconv.Atoi(line4[3])

		// 0, 1, 2, 3 => left, up, right, down

		var vertical bool
		var reverse bool
		switch direction {
		case 0:
			vertical = false
			reverse = false
		case 1:
			vertical = true
			reverse = false
		case 2:
			vertical = false
			reverse = true
		case 3:
			vertical = true
			reverse = true
		}

		// BEFORE
		// for _, row := range board {
		// 	fmt.Println(row)
		// }
		// fmt.Println()

		resolve(vertical, reverse, &board)

		// AFTER
		for r := 0; r < 4; r++ {
			for c := 0; c < 4; c++ {
				fmt.Printf("%d ", board[r][c])
			}
			fmt.Println()
		}
	}
}

func resolve(vertical bool, reverse bool, table *[4][4]int) {

	var columns []int
	if reverse {
		columns = []int{3, 2, 1, 0}
	} else {
		columns = []int{0, 1, 2, 3}
	}

	var line, column, coordA, coordB int
	//
	var nextAfterRef, nextAfterRefA, nextAfterRefB int

	for i := 0; i < 4; i++ {
		aux := columns[0]
		endCmp := false

	OUTER:
		for _, j := range columns { // GET COLUMN INDEX TO START
			//
			for {
				if endCmp {
					break OUTER
				}

				if reverse {
					aux--
				} else {
					aux++
				}

				if aux == 0 || aux == 3 {
					endCmp = true
				}

				if vertical {

					line = j
					column = i

					coordA = aux
					coordB = column

					nextAfterRef = line + 1
					if reverse {
						nextAfterRef = line - 1
					}

					nextAfterRefA = nextAfterRef
					nextAfterRefB = column

				} else {

					line = i
					column = j

					coordA = line
					coordB = aux

					nextAfterRef = column + 1
					if reverse {
						nextAfterRef = column - 1
					}
					nextAfterRefA = line
					nextAfterRefB = nextAfterRef
				}

				current := table[line][column]
				next := table[coordA][coordB]

				if current == 0 {

					if next != 0 {
						// current GETS next and next GETS 0
						table[line][column] = table[coordA][coordB]
						table[coordA][coordB] = 0
					}

				} else if current == next {

					// current GETS double value (current + next) and next GETS 0
					table[line][column] = current + next
					table[coordA][coordB] = 0
					break

				} else if current != next && next != 0 {

					// check if has to move
					if aux != nextAfterRef {

						// move value for the next position close to the current position
						// and puts 0 in the compared position
						table[nextAfterRefA][nextAfterRefB] = next
						table[coordA][coordB] = 0

					}
					break
				}
			}
		}
	}
}
