package main

import "fmt"

// ListNode for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	list1 := &ListNode{
		Val:  5,
		Next: nil,
		// Next: &ListNode{
		// Val: 4,
		// Next: &ListNode{
		// Val:  3,
		// Next: nil,
		// },
		// },
	}

	list2 := &ListNode{
		Val:  5,
		Next: nil,
		// Next: &ListNode{
		// 	Val: 6,
		// 	Next: &ListNode{
		// 		Val:  4,
		// 		Next: nil,
		// 	},
		// },
	}

	addTwoNumbers(list1, list2)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	listFinal := &ListNode{
		Val: (l1.Val + l2.Val) % 10,
	}

	anteriorList := listFinal

	nextList1 := l1
	nextList2 := l2

	sumRest := (l1.Val + l2.Val) / 10

	if nextList1.Next == nil && nextList2.Next == nil && sumRest == 0 {
		return listFinal
	}

	nextList1 = nextList1.Next
	nextList2 = nextList2.Next

	for {
		v1 := 0
		v2 := 0

		if nextList1 == nil && nextList2 == nil && sumRest == 0 {
			break
		}

		if nextList1 != nil {
			v1 = nextList1.Val
		}

		if nextList2 != nil {
			v2 = nextList2.Val
		}

		fmt.Println(v1, v2)

		sum := sumRest
		sum += v1 + v2

		sumRest = sum / 10

		sum = sum % 10

		list := &ListNode{
			Val: sum,
		}

		anteriorList.Next = list

		anteriorList = list

		if nextList1 != nil {
			nextList1 = nextList1.Next
		}

		if nextList2 != nil {
			nextList2 = nextList2.Next
		}
	}

	return listFinal
}
