package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

//递归
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Swap Nodes in Pairs.
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head.Next
	head.Next = swapPairs(head.Next.Next)
	dummy.Next.Next = head

	return dummy.Next
}

//非递归
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Swap Nodes in Pairs.
//func swapPairs(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	now := head.Next
//	ret := now
//	head.Next = now.Next
//	now.Next = head
//	for head.Next != nil && head.Next.Next != nil {
//		now = head.Next
//		head.Next = now.Next
//		after := now.Next
//		now.Next = after.Next
//		after.Next = now
//		head = now
//	}
//	return ret
//}

func main() {
	five := &ListNode{Val: 5}
	four := &ListNode{Val: 4, Next: five}
	three := &ListNode{Val: 3, Next: four}
	two := &ListNode{Val: 2, Next: three}
	one := &ListNode{Val: 1, Next: two}
	node := swapPairs(one)
	for {
		fmt.Printf("%v ->", node.Val)
		if node.Next == nil {
			break
		}
		node = node.Next
	}
}
