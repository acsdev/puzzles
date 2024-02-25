package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Stack struct {
	Data []int
}

func (s *Stack) push(value int) {
	s.Data = append(s.Data, value)
}

func (s *Stack) pushAll(value []int) {
	s.Data = append(s.Data, value...)
}

func (s *Stack) peek() int {
	if s.empty() {
		return -1
	}
	size := len(s.Data) - 1
	return s.Data[size]
}

func (s *Stack) pop() int {
	size := len(s.Data) - 1
	item := s.Data[size]
	s.Data = s.Data[:size]
	return item
}

func (s *Stack) empty() bool {
	return len(s.Data) == 0
}

func (s *Stack) size() int {
	return len(s.Data)
}

func (s *Stack) clean() {
	s.Data = make([]int, 0)
}

type Player struct {
	ID   int
	Deck *Stack
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

	for {
		scanner.Scan()
		line1 := scanner.Text()
		if line1 == "0 0" {
			break
		}

		scanner.Scan()
		line2 := scanner.Text()
		// reading input data
		input1 := strings.Split(line1, " ")
		input2 := strings.Split(line2, " ")

		deckAmount, _ := strconv.Atoi(input1[0])
		numerOfPlayers, _ := strconv.Atoi(input1[1])
		mainDeck := &Stack{
			Data: make([]int, deckAmount),
		}
		for index, card := range input2 {
			mainDeck.Data[index], _ = strconv.Atoi(card)
		}

		players := make([]*Player, numerOfPlayers)
		for indexPlayer := 0; indexPlayer < numerOfPlayers; indexPlayer++ {
			players[indexPlayer] = &Player{
				ID: indexPlayer + 1,
				Deck: &Stack{
					Data: make([]int, 0),
				},
			}
		}
		play(mainDeck, players)
	}
}

func play(mainDeck *Stack, players []*Player) {

	// turn deck upside down
	i := 0
	j := len(mainDeck.Data) - 1
	for i < j {
		mainDeck.Data[i], mainDeck.Data[j] = mainDeck.Data[j], mainDeck.Data[i]
		i = i + 1
		j = j - 1
	}

	playerTurn := 0
	discartedCards := make([]int, 0)
OUTER:
	for {
		if mainDeck.empty() {
			break
		}

		player := players[playerTurn]
		cardTurn := mainDeck.peek()

		has, card, discarted := handleCardInDiscartedCard(cardTurn, discartedCards)
		if has {
			discartedCards = discarted
			player.Deck.push(card)
			player.Deck.push(mainDeck.pop())
			continue // get another card
		}

		for _, anotherPlayer := range players {
			cardOnTheTop := anotherPlayer.Deck.peek()
			if cardOnTheTop == cardTurn {
				if player.ID != anotherPlayer.ID {
					// means that the card is equal to the card on top of player that is not the current player
					// so, the player can steel their deck
					player.Deck.pushAll(anotherPlayer.Deck.Data)
					anotherPlayer.Deck.clean()
				}
				player.Deck.push(mainDeck.pop())
				continue OUTER
			}
		}

		discartedCards = append(discartedCards, mainDeck.pop())
		if playerTurn == len(players)-1 {
			playerTurn = 0
			continue
		}
		playerTurn++
	}

	// print winners
	winersMax := -1
	winersID := make([]int, 0)
	for _, player := range players {
		id := player.ID
		size := player.Deck.size()
		if size >= winersMax {
			if size == winersMax {
				// means that were previus winner
				winersID = append(winersID, id)
			} else {
				winersID = make([]int, 0)
				winersID = append(winersID, id)
			}
			winersMax = size
		}
	}
	sort.Ints(winersID)
	fmt.Print(winersMax)
	if len(winersID) > 0 {
		for _, v := range winersID {
			fmt.Print(" ")
			fmt.Print(v)
		}
	} else {
		fmt.Print(" ")
		fmt.Print(winersID[0])
	}
	fmt.Println()
}

func handleCardInDiscartedCard(card int, discarted []int) (bool, int, []int) {
	for indexCardOnDeck, cardOnDeck := range discarted {
		if cardOnDeck == card {
			discarted = append(discarted[:indexCardOnDeck], discarted[indexCardOnDeck+1:]...) // remove  item from discated
			return true, cardOnDeck, discarted
		}
	}
	return false, -1, nil
}
