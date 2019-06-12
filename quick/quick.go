package main

import (
	"fmt"
	"time"
)

type quick struct {
	arr []int
}

func main() {
	q := &quick{
		arr: []int{10, 12, 90, 5, 6},
	}

	start := time.Now()
	q.quick(0, len(q.arr)-1)
	end := time.Since(start)
	fmt.Println("time: ", end, "result: ", q.arr)

}

func (q *quick) quick(l, h int) {
	fmt.Println("lh", l, h)
	if l < h {
		j := q.swap(l, h)
		q.quick(l, j-1)
		q.quick(j+1, h)
	}
}

func (q *quick) swap(l, h int) int {
	pivot := q.arr[l]
	j := h
	i := l

	for i < j {
		if q.arr[i] < pivot {
			i++
		}
		if q.arr[j] > pivot {
			j--
		}

		aux := q.arr[i]
		q.arr[i] = q.arr[j]
		q.arr[j] = aux
	}
	return j
}
