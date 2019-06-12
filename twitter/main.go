/*
We have a list of letters with lowercase letters inside and a list of words.

Write a function numSets which will take 2 arguments:
  1. letters: []rune list of letters, such as []rune{'i', 'n', 't', 'e', 'r', 'v', 'i', 'e', 'w'}
  2. words: []string list of words, such as: words = []string{"view", "rent"}

Return how many "letters lists" you would need to build those words.
If it’s not possible to build the words from the list, return -1.

Ex.
  numSets([]rune{'a', 'b', 'c', 'd'}, []string{"a"}) -> 1
  numSets([]rune{'a', 'b', 'b', 'c'}, []string{"abc"}) -> 1
  numSets([]rune{'a', 'b', 'c', 'd'}, []string{"ab", "ca"}) -> 2
  numSets([]rune{'a', 'b', 'c', 'd'}, []string{"abc", "da", "bca"}) -> 3
  numSets([]rune{'a', 'b', 'c', 'd'}, []string{"e"}) -> -1
  numSets([]rune{'i', 'n', 't', 'e', 'r', 'v', 'i', 'e', 'w'}, []string{"view", "rent"}) -> 1
  numSets([]rune{'i', 'n', 't', 'e', 'r', 'v', 'i', 'e', 'w'}, []string{"view", "view"}) -> 2
*/

package main

import (
	"fmt"
	"sort"
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
		strFreq[fmt.Sprintf("%c", s)]++
	}

	sf := make([]strsFreq, len(strFreq))

	i := 0
	for str, freq := range strFreq {
		sf[i].str = str
		sf[i].freq = freq
		i++
	}

	sort.Slice(sf, func(i, j int) bool {
		return sf[i].freq > sf[j].freq
	})

	for _, sf := range sf {
		if runeFreq[sf.str] == 0 {
			return -1
		}
	}

	if sf[0].freq > runeFreq[sf[0].str] {
		sets += sf[0].freq - runeFreq[sf[0].str]
	}

	return sets
}
