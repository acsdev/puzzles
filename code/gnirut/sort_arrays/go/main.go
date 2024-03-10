package main

import (
	"fmt"
	"sort"
)

func main() {

	array := [9]int{-1, 1, -6, 4, 5, -6, 1, 4, 1}

	valueByOcurrences := make(map[int]int, 0)
	// group value by ocurrences
	for _, v := range array {

		value, has := valueByOcurrences[v]
		if !has {
			value = 1
		} else {
			value += 1
		}
		valueByOcurrences[v] = value
	}

	// group occurences by values
	occurences := make(map[int][]int, 0)
	for k, v := range valueByOcurrences {

		var aux []int
		if len(occurences) == 0 || len(occurences[v]) == 0 {
			aux = make([]int, 0)
		} else {
			aux = occurences[v]
		}
		aux = append(aux, k)
		occurences[v] = aux
	}

	// reorder array with items
	index := 0
	for k := 1; k < len(occurences)+1; k++ {
		v := occurences[k]
		sort.Ints(v) // read in reserse mode
		//
		for i := len(v) - 1; i >= 0; i-- {
			for j := 0; j < k; j++ {
				array[index] = v[i]
				index++
			}
		}
	}

	fmt.Println(array)
}
