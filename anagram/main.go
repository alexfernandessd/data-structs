package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	test := []string{
		"eatb",
		// "teab",
		// "nad",
		// "dan",
	}

	starts := time.Now()
	groupStrings := solution(test)
	ends := time.Since(starts)
	fmt.Println(groupStrings, ends)
}

func solution(strs []string) [][]string {

	stringMap := map[string][]string{}

	for _, str := range strs {
		ss := strings.Split(str, "")
		strs := quickSort(ss, 0, len(ss)-1)
		s := strings.Join(strs, "")
		stringMap[s] = append(stringMap[s], str)
	}

	arrayFinal := make([][]string, len(stringMap))

	j := 0
	for _, strs := range stringMap {
		arrayFinal[j] = strs
		j++
	}

	return arrayFinal
}

func quickSort(strings []string, l, h int) []string {
	strs := strings
	if l < h {
		strs, j := swap(strs, l, h)
		strs = quickSort(strs, l, j-1)
		strs = quickSort(strs, j+1, h)
	}

	return strs
}

func swap(strs []string, l, h int) ([]string, int) {

	pivot := strs[l]
	i := l
	j := h

	for i < j {
		if strs[j] > pivot {
			j--
		}
		if strs[i] < pivot {
			i++
		}

		aux := strs[i]
		strs[i] = strs[j]
		strs[j] = aux
	}

	return strs, j
}
