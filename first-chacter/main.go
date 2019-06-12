package main

import (
	"fmt"
)

func main() {
	s := "loveleetcode"

	i := bla(s)

	fmt.Println(i)

}

func bla(s string) int {
	strs := map[string]int{}

	for _, str := range s {
		strs[fmt.Sprintf("%c", str)]++
	}

	for i, str := range s {
		if strs[fmt.Sprintf("%c", str)] == 1 {
			return i
		}
	}

	return -1
}
