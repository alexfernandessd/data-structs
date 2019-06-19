package main

import "fmt"

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	list := &ListNode{
		Val:  1,
		Next: nil,
		// Next: &ListNode{
		// 	Val: 2,
		// 	// Next: nil,
		// 	Next: &ListNode{
		// 		Val: 3,
		// 		Next: &ListNode{
		// 			Val: 4,
		// 			Next: &ListNode{
		// 				Val:  5,
		// 				// Next: nil,
		// 				Next: &ListNode{
		// 					Val:  6,
		// 					Next: nil,
		// 				},
		// 			},
		// 		},
		// 	},
		// },
	}

	reverseKGroup(nil, 2)
	for {
		if list == nil {
			break
		}
		fmt.Print(list.Val)
		list = list.Next
	}
}

func reverseKGroup(head *ListNode, k int) {
	c := 0
	nums := make([]int, k)
	atualList := head

	times := 0

	for {
		nums[c] = atualList.Val

		if c == k-1 {
			listAtual := head
			j := 1
			for i := len(nums) - 1; i >= 0; {
				if listAtual == nil {
					break
				}
				if j > times {
					listAtual.Val = nums[i]
					i--
				}
				listAtual = listAtual.Next
				j++
			}
			c = 0
			times += k
			if atualList.Next == nil {
				break
			}
			atualList = atualList.Next
			continue
		}

		c++
		if atualList.Next == nil {
			break
		}
		atualList = atualList.Next
	}
}
