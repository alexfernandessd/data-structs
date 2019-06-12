package main

import (
	"fmt"
	"time"
)

func main() {

	arr := createHeap()
	fmt.Println("heap: ", arr)

	arr2 := heapify()
	fmt.Println("heapify: ", arr2)
}

func heapify() []int {
	start := time.Now()
	arr := []int{10, 20, 50, 80, 900, 1000, 5, 6}

	for i := len(arr) - 1; i >= 0; i-- {
		childL := i + 1*2
		childR := i + 1*2 + 1

		maxChild := childL

		if childL <= len(arr) {
			if childR <= len(arr) && arr[childR-1] > arr[childL-1] {
				maxChild = childR
			}
			if arr[maxChild-1] > arr[i] {
				aux := arr[i]
				arr[i] = arr[maxChild-1]
				arr[maxChild-1] = aux

				j := i
				for parent := int(i / 2); parent < j; parent = int(parent / 2) {
					if arr[j] > arr[parent] {
						aux := arr[parent]
						arr[parent] = arr[j]
						arr[j] = aux
						j = parent
						continue
					}
					break
				}

			}
		}
	}

	end := time.Since(start)

	fmt.Println("time: ", end)

	return arr
}

func createHeap() []int {
	start := time.Now()
	arr := []int{10, 20, 50, 80, 900, 1000, 5, 6}

	for i, a := range arr {
		childL := i + 1*2
		childR := i + 1*2 + 1

		maxChild := childL

		if childL <= len(arr) {
			if childR <= len(arr) && arr[childR-1] > arr[childL-1] {
				maxChild = childR
			}
			if arr[maxChild-1] > arr[i] {
				arr[i] = arr[maxChild-1]
				arr[maxChild-1] = a

				j := i

				for parent := int(i / 2); parent < j; parent = int(parent / 2) {
					if arr[j] > arr[parent] {
						aux := arr[parent]
						arr[parent] = arr[j]
						arr[j] = aux
						j = parent
						continue
					}
					break
				}

			}
		}

	}

	end := time.Since(start)
	fmt.Println("time: ", end)

	return arr
}
