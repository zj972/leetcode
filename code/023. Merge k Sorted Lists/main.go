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

//优化list遍历过程
//Runtime: 12 ms, faster than 93.33% of Go online submissions for Merge k Sorted Lists.
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	interval := 1
	for interval < len(lists) {
		for i := 0; i < len(lists)-interval; i += interval * 2 {
			lists[i] = mergeTwoLists(lists[i], lists[i+interval])
		}
		interval *= 2
	}

	return lists[0]
}

//Runtime: 216 ms, faster than 29.33% of Go online submissions for Merge k Sorted Lists.
//func mergeKLists(lists []*ListNode) *ListNode {
//	if len(lists) == 0 {
//		return nil
//	}
//	if len(lists) == 1 {
//		return lists[0]
//	}
//	lists[1] = mergeTwoLists(lists[0], lists[1])
//	return mergeKLists(lists[1:])
//}

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
	five2 := &ListNode{Val: 5}
	four2 := &ListNode{Val: 4, Next: five2}
	three2 := &ListNode{Val: 3, Next: four2}
	two2 := &ListNode{Val: 2, Next: three2}
	one2 := &ListNode{Val: 1, Next: two2}
	node := mergeKLists([]*ListNode{one, one1, one2})
	for {
		fmt.Printf("%v ->", node.Val)
		if node.Next == nil {
			break
		}
		node = node.Next
	}
}
