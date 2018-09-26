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

//Runtime: 0 ms
//Your runtime beats 100.00 % of golang submissions.
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var getNextNum func(node *ListNode) int
	getNextNum = func(node *ListNode) int {
		if node.Next != nil {
			return getNextNum(node.Next) + 1
		}
		return 1
	}
	nowNode := head
	num := getNextNum(nowNode)
	if num == n {
		return nowNode.Next
	}
	for i := 0; i < num-n-1; i++ {
		nowNode = nowNode.Next
	}
	nowNode.Next = nowNode.Next.Next
	return head
}

func main() {
	five := &ListNode{Val: 5}
	four := &ListNode{Val: 4, Next: five}
	three := &ListNode{Val: 3, Next: four}
	two := &ListNode{Val: 2, Next: three}
	one := &ListNode{Val: 1, Next: two}
	node := removeNthFromEnd(one, 2)
	for {
		fmt.Printf("%v ->", node.Val)
		if node.Next == nil {
			break
		}
		node = node.Next
	}
}
