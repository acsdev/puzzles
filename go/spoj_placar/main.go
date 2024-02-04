package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Student struct {
	Name  string
	Score int
}

// OrderByPointsAndName implements sort.Interface for []Student based on
type OrderByScoreAndName []Student

func (o OrderByScoreAndName) Len() int {
	return len(o)
}
func (o OrderByScoreAndName) Less(i, j int) bool {
	if o[i].Score == o[j].Score {
		return o[i].Name < o[j].Name
	}
	return o[i].Score > o[j].Score
}
func (o OrderByScoreAndName) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

func getIntValue(value string) int {
	converted, err := strconv.Atoi(value)
	if err != nil {
		log.Panic(err)
	}
	return converted
}

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

	instance := 0
	for {
		scanner.Scan()
		row := scanner.Text()
		if row == "" {
			break
		}

		numberOfStudents := getIntValue(row)
		studants := make([]Student, numberOfStudents)
		for index := range studants {
			scanner.Scan()
			data := strings.Split(scanner.Text(), " ")
			studant := Student{
				Name:  data[0],
				Score: getIntValue(data[1]),
			}
			studants[index] = studant
		}

		sort.Sort(OrderByScoreAndName(studants))

		instance += 1
		fmt.Printf("Instancia %d\n", instance)
		fmt.Printf("%s\n", studants[len(studants)-1].Name)
	}
}
