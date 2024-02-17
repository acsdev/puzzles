package main

import (
	"bufio"
	"fmt"
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

	for scanner.Scan() {
		value, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		if value == 42 {
			break
		}
		fmt.Println(value)
	}
}
