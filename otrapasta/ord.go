package main

import "fmt"

func main() {
	arr := []int{50, 40, 30, 10, 20}
	// arrFinal := bubble(arr)
	// fmt.Println(arrFinal)

	arrFinal2 := selection(arr)
	fmt.Println(arrFinal2)
}

func bubble(arr []int) []int {
	for i := range arr {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				aux := arr[i]
				arr[i] = arr[j]
				arr[j] = aux
			}
		}
	}
	return arr
}

func selection(arr []int) []int {
	for i := range arr {
		m := 0
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[m] {
				m = j
			}
		}

		last := arr[len(arr)-i-1]
		arr[len(arr)-i-1] = arr[m]
		arr[m] = last
	}

	return arr
}
