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

//Runtime: 8 ms, faster than 32.79% of Go online submissions for Reverse Nodes in k-Group.
//func reverseKGroup(head *ListNode, k int) *ListNode {
//	l := length(head)
//	if l < k {
//		return head
//	}
//	dummp := new(ListNode)
//	dummp.Next = head
//	b, a := swap(dummp, head, k)
//	ret := b.Next
//	l -= k
//	for l >= k {
//		_, a = swap(a, a.Next, k)
//		l -= k
//	}
//	return ret
//}
//
//func length(head *ListNode) (l int) {
//	if head == nil {
//		return 0
//	}
//	if head.Next == nil {
//		return 1
//	}
//	return length(head.Next) + 1
//}
//
//func swap(before *ListNode, after *ListNode, k int) (b *ListNode, a *ListNode) {
//	if k == 1 {
//		return before, after
//	}
//	mid := before.Next
//	before.Next = after.Next
//	after.Next = before.Next.Next
//	before.Next.Next = mid
//	return swap(before, after, k-1)
//}

//Runtime: 4 ms, faster than 100.00% of Go online submissions for Reverse Nodes in k-Group.
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k < 2 {
		return head
	}
	return helper(head, k)
}
func helper(head *ListNode, k int) *ListNode {
	tail, needReverse := getTail(head, k)

	if needReverse {
		tailNext := tail.Next
		tail.Next = nil
		head, tail = reverse(head, tail)
		tail.Next = helper(tailNext, k)
	}

	return head
}

func getTail(head *ListNode, k int) (*ListNode, bool) {
	for k > 1 && head != nil {
		k--
		head = head.Next
	}
	return head, k == 1 && head != nil
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	curPre, cur := head, head.Next
	for cur != nil {
		curPre, cur, cur.Next = cur, cur.Next, curPre
	}
	return tail, head
}

func main() {
	five := &ListNode{Val: 5}
	four := &ListNode{Val: 4, Next: five}
	three := &ListNode{Val: 3, Next: four}
	two := &ListNode{Val: 2, Next: three}
	one := &ListNode{Val: 1, Next: two}
	_ = one
	node := reverseKGroup(nil, 2)
	fmt.Println(node)
	for {
		fmt.Printf("%v ->", node.Val)
		if node.Next == nil {
			break
		}
		node = node.Next
	}
}
