/*
1.避免生成不必要的节点
2.在同一个func中尽量避免相同条件多次调用
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	now := head
	carry := 0
	for {
		x := 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		y := 0
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum := x + y + carry
		carry = sum / 10
		now.Val = sum % 10
		if l1 == nil && l2 == nil {
			break
		}
		now.Next = &ListNode{}
		now = now.Next
	}
	if carry > 0 {
		now.Next = &ListNode{Val: carry}
	}
	return head
}

func main() {
	c := &ListNode{3, nil}
	b := &ListNode{2, c}
	a := &ListNode{1, b}
	y := &ListNode{8, nil}
	x := &ListNode{7, y}
	//321+87=1308
	res := addTwoNumbers(a, x)
	for res != nil {
		fmt.Print(res.Val)
		res = res.Next
	}
}

// wrong answer by 刘乾
//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	out := new(ListNode)
//	out.Val = 0
//	cur := out
//	carry := 0
//	if l1 == nil {
//		out = l2
//		return out
//	}
//	if l2 == nil {
//		out = l1
//		return out
//	}
//	if l1 == nil && l2 == nil {
//		return out
//	}
//	for l1.Next != nil || l2.Next != nil {
//		fmt.Println(l1, l2)
//		x, y := 0, 0
//		if l1 != nil {
//			x = l1.Val
//		} else {
//			x = 0
//		}
//		if l2 != nil {
//			y = l2.Val
//		} else {
//			y = 0
//		}
//		sum := x + y + carry
//		carry = sum / 10
//		cur.Next = &ListNode{Val: sum % 10}
//		cur = cur.Next
//		if l1 != nil {
//			l1 = l1.Next
//		}
//		if l2 != nil {
//			l2 = l2.Next
//		}
//	}
//	if (carry) > 0 {
//		cur.Next = &ListNode{Val: carry}
//	}
//	return out.Next
//}
