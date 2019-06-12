package main

func main1() {

	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 8,
			// 	Next: &ListNode{
			// 		Val:  3,
			// 		Next: nil,
			// 	},
		},
		// Next: nil,
	}

	list2 := &ListNode{
		Val: 0,
		// Next: &ListNode{
		// 	Val: 6,
		// 	Next: &ListNode{
		// 		Val:  4,
		// 		Next: nil,
		// 	},
		// },
		Next: nil,
	}

	addTwoNumbers1(list1, list2)
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {

	lista1 := []int{}

	next := true
	nextList := l1

	for next != false {
		lista1 = append(lista1, nextList.Val)

		if nextList.Next != nil {
			nextList = nextList.Next
		} else {
			next = false
		}
	}

	lista2 := []int{}

	next = true
	nextList = l2

	for next != false {
		lista2 = append(lista2, nextList.Val)

		if nextList.Next != nil {
			nextList = nextList.Next
		} else {
			next = false
		}
	}

	i1 := 0
	i2 := 0

	listFinal := &ListNode{
		Val: (lista1[i1] + lista2[i2]) % 10,
	}

	sumRest := (lista1[i1] + lista2[i2]) / 10

	anterior := listFinal

	var n1 int
	var n2 int
	var sum int

	i1++
	i2++

	for {

		if i1 >= len(lista1) {
			n1 = 0
		} else {
			n1 = lista1[i1]
		}

		if i2 >= len(lista2) {
			n2 = 0
		} else {
			n2 = lista2[i2]
		}

		if i1 >= len(lista1) && i2 >= len(lista2) && sumRest == 0 {
			break
		}

		sum = sumRest
		sum += n1 + n2

		sumRest = sum / 10

		sum = sum % 10

		list := &ListNode{
			Val: sum,
		}

		anterior.Next = list

		anterior = list

		i1++
		i2++

	}

	return listFinal

}
