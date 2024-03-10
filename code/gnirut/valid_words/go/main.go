package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "cb"
	k := [5]string{"cd", "bd", "cccb", "bcc", "bcdcb"}

	count := 0
OUTER:
	for _, value := range k {

		chars := strings.Split(value, "")

		for _, char := range chars {
			if !strings.Contains(s, char) {
				continue OUTER
			}
		}
		count++
	}

	fmt.Println(count)
}
