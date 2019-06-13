package main

import (
	"fmt"
	"strings"
)

type strsFreq struct {
	str  string
	freq int
}

func main() {
	runes := []rune{'i', 'n', 't', 'e', 'r', 'v', 'i', 'e', 'w'}
	arrStrings := []string{"view", "view"}

	sets := numSets(runes, arrStrings)
	fmt.Println("result: ", sets)
}

func numSets(runes []rune, arrStrings []string) int {
	sets := 1

	runeFreq := map[string]int{}

	for _, r := range runes {
		runeFreq[fmt.Sprintf("%c", r)]++
	}

	strFreq := map[string]int{}

	strConcat := strings.Join(arrStrings, "")

	for _, s := range strConcat {
		if runeFreq[fmt.Sprintf("%c", s)] == 0 {
			return -1
		}
		strFreq[fmt.Sprintf("%c", s)]++
	}

	sf := make([]strsFreq, len(strFreq))

	i := 0
	for str, freq := range strFreq {
		sf[i].str = str
		sf[i].freq = freq
		i++
	}

	sf = maxHeapSort(sf)

	if sf[0].freq > runeFreq[sf[0].str] {
		sets += sf[0].freq - runeFreq[sf[0].str]
	}

	return sets
}

func maxHeapSort(arr []strsFreq) []strsFreq {
	for i := len(arr) - 1; i >= 0; i-- {
		parent := i / 2
		if parent >= 0 && arr[i].freq > arr[parent].freq {
			auxParent := arr[parent]
			arr[parent] = arr[i]
			arr[i] = auxParent
		}
	}

	return arr
}
