package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Player struct {
	Name              string
	Points            []int
	TotalPoint        int
	MaxPointDiscarted int
	MinPointDiscarted int
}

// OrderByPointsAndName implements sort.Interface for []Player based on
type OrderByPointsAndName []Player

func (o OrderByPointsAndName) Len() int {
	return len(o)
}
func (o OrderByPointsAndName) Less(i, j int) bool {
	if o[i].TotalPoint == o[j].TotalPoint {
		return o[i].Name < o[j].Name
	}
	return o[i].TotalPoint > o[j].TotalPoint
}
func (o OrderByPointsAndName) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
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

	cases := make([][]Player, 0)
	for {
		scanner.Scan()
		row := scanner.Text()
		if row == "0" {
			break
		}

		numberOfPlayers, _ := strconv.Atoi(row)
		players := make([]Player, numberOfPlayers)
		for index := range players {
			scanner.Scan()
			firstLine := scanner.Text()
			scanner.Scan()
			secondLine := scanner.Text()

			player := Player{}
			player.Name = firstLine
			player.Points = []int{}
			player.MinPointDiscarted = 9999
			player.MaxPointDiscarted = 0
			points := strings.Split(secondLine, " ")
			for _, point := range points {
				pointAsInt, _ := strconv.Atoi(point)
				if player.MinPointDiscarted > pointAsInt {
					player.MinPointDiscarted = pointAsInt
				}
				if player.MaxPointDiscarted < pointAsInt {
					player.MaxPointDiscarted = pointAsInt
				}
				player.TotalPoint += pointAsInt
				player.Points = append(player.Points, pointAsInt)
			}
			player.TotalPoint -= player.MinPointDiscarted
			player.TotalPoint -= player.MaxPointDiscarted
			players[index] = player
		}

		sort.Sort((OrderByPointsAndName(players))) // sort by point and name
		cases = append(cases, players)
		//
		// PrintPlayers(players)
	}

	for index, players := range cases {
		fmt.Println("Teste " + strconv.Itoa(index+1))
		rank := 1
		for index, player := range players {
			row := index + 1
			if index == 0 {
				fmt.Printf("%d %d %s \n", rank, player.TotalPoint, player.Name)
			} else {
				if players[index-1].TotalPoint != player.TotalPoint {
					rank = row
				}
				fmt.Printf("%d %d %s \n", rank, player.TotalPoint, player.Name)
			}
		}
	}
}

func PrintPlayers(players []Player) {
	for _, player := range players {
		fmt.Println(player.Name)
		fmt.Println(player.Points)
		fmt.Println(player.TotalPoint)
		fmt.Println(player.MinPointDiscarted)
		fmt.Println(player.MaxPointDiscarted)
	}
}
