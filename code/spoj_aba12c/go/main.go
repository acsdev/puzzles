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
	scenarios, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < scenarios; i++ {

		scanner.Scan()
		firstLine := strings.Split(scanner.Text(), " ")

		amountKg, err := strconv.Atoi(firstLine[1]) // amountKg
		if err != nil {
			amountKg = 0
		}

		scanner.Scan()
		secondLine := strings.Split(scanner.Text(), " ")

		packageInKgList := make([]int, 0)
		packageInKgPrice := make(map[int]int)
		for index, value := range secondLine {
			_price, err := strconv.Atoi(value)
			if err != nil {
				_price = 0
			}

			if _price < 0 {
				continue
			}
			packageInKg := index + 1
			packageInKgList = append(packageInKgList, packageInKg)
			packageInKgPrice[packageInKg] = _price
		}

		memo := make(map[int]int)
		//
		fmt.Println(solution(amountKg, packageInKgList, packageInKgPrice, memo))
	}
}

func solution(amoutKg int, kgList []int, kgPrice map[int]int, memo map[int]int) int {

	if amoutKg == 0 {
		return 0
	}

	if amoutKg < 0 {
		return -1
	}

	memoizedPrice, has := memo[amoutKg]
	if has {
		return memoizedPrice
	}

	minPrice := -1
	for _, kg := range kgList {
		subAmountKg := amoutKg - kg
		subPrice := solution(subAmountKg, kgList, kgPrice, memo)
		if subPrice >= 0 {
			price := kgPrice[kg]
			combinedPrice := price + subPrice
			if combinedPrice < minPrice || minPrice == -1 {
				minPrice = combinedPrice
			}
		}
	}

	memo[amoutKg] = minPrice
	return minPrice
}
