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

//Runtime: 12 ms, faster than 0.47% of Go online submissions for Merge Two Sorted Lists.
//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	var node = new(ListNode)
//	ret := node
//	for l1 != nil || l2 != nil {
//		if l1 == nil {
//			node.Next = l2
//			break
//		}
//		if l2 == nil {
//			node.Next = l1
//			break
//		}
//		fmt.Println(l1.Val, l2.Val)
//		if l1.Val < l2.Val {
//			node.Next = l1
//			l1 = l1.Next
//		} else {
//			node.Next = l2
//			l2 = l2.Next
//		}
//		node = node.Next
//	}
//	return ret.Next
//}

//Runtime: 4 ms, faster than 100.00% of Go online submissions for Merge Two Sorted Lists.
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	p := l1
	q := l2

	var head *ListNode
	var current *ListNode

	for p != nil || q != nil {
		var next *ListNode
		if p == nil {
			next = q
			q = q.Next
		} else if q == nil {
			next = p
			p = p.Next
		} else if p.Val < q.Val {
			next = p
			p = p.Next
		} else {
			next = q
			q = q.Next
		}

		if head == nil {
			head = next
		} else {
			current.Next = next
		}
		current = next
	}
	return head
}

func main() {
	five := &ListNode{Val: 5}
	four := &ListNode{Val: 4, Next: five}
	three := &ListNode{Val: 3, Next: four}
	two := &ListNode{Val: 2, Next: three}
	one := &ListNode{Val: 1, Next: two}
	five1 := &ListNode{Val: 5}
	four1 := &ListNode{Val: 4, Next: five1}
	three1 := &ListNode{Val: 3, Next: four1}
	two1 := &ListNode{Val: 2, Next: three1}
	one1 := &ListNode{Val: 1, Next: two1}
	node := mergeTwoLists(one, one1)
	for {
		fmt.Printf("%v ->", node.Val)
		if node.Next == nil {
			break
		}
		node = node.Next
	}
}
