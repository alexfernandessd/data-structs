package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 1, 1, 1, 2, 2, 3, 3}
	arrFinal := barcode(arr)
	fmt.Println(arrFinal)
}

type array struct {
	num  int
	freq int
}

func barcode(arr []int) []int {
	arrFinal := []int{}

	arrMap := map[int]int{}

	for _, num := range arr {
		arrMap[num]++
	}
	array := make([]array, len(arrMap))

	i := 0
	for num, freq := range arrMap {
		array[i].freq = freq
		array[i].num = num
		i++
	}

	lastNumber := -1
	j := 0
	for len(arrFinal) != len(arr) {
		sort.Slice(array, func(i, j int) bool {
			return array[i].freq > array[j].freq
		})

		if array[j].num != lastNumber {
			lastNumber = array[j].num
			arrFinal = append(arrFinal, array[j].num)
			array[j].freq--
			if array[j].freq == 0 {
				array = append(array[:j], array[j+1:]...)
			}
			j = 0
		} else {
			if j >= len(array)-1 {
				j = 0
			} else {
				j++
			}
		}
	}

	return arrFinal
}
