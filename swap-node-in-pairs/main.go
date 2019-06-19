package main

import "fmt"

// ListNode for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					// Next: nil,
				},
			},
		},
	}

	listFinal := swapNodes(list)

	for listFinal != nil {
		fmt.Println(listFinal.Val)
		listFinal = listFinal.Next
	}
}

func swapNodes(list *ListNode) *ListNode {
	prevList := list

	for prevList != nil {
		auxVal := prevList.Val
		atualList := prevList.Next
		if atualList != nil {
			prevList.Val = atualList.Val
			atualList.Val = auxVal
			prevList = atualList.Next
		} else {
			break
		}
	}

	return list
}
