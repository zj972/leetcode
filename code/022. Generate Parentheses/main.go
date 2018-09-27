package main

import (
	"fmt"
)

//Runtime: 8 ms, faster than 100.00% of Go online submissions for Generate Parentheses.
//type node struct {
//	item string
//	num  int
//}
//
//func generateParenthesis(n int) []string {
//	listNode := []*node{}
//	list := []string{}
//	listNode = add(&node{}, n)
//	for j := 0; j < len(listNode); j++ {
//		if listNode[j].num == 0 {
//			list = append(list, listNode[j].item)
//		} else {
//			for k := 0; k < listNode[j].num; k++ {
//				listNode[j].item = listNode[j].item + ")"
//			}
//			list = append(list, listNode[j].item)
//		}
//	}
//	return list
//}
//
//func add(l *node, n int) (o []*node) {
//	if n == 0 {
//		if l.num != 0 {
//			for i := 0; i < l.num; i++ {
//				l.item += ")"
//			}
//			l.num = 0
//		}
//		return []*node{l}
//	}
//	if l.num == 0 {
//		o = add(&node{l.item + "(", 1}, n-1)
//	} else {
//		o = add(&node{l.item + "(", l.num + 1}, n-1)
//		o = append(o, add(&node{l.item + ")", l.num - 1}, n)...)
//	}
//	return o
//}

//存在冗余数据（长度小于max*2的数据）
//Runtime: 8 ms, faster than 100.00% of Go online submissions for Generate Parentheses.
func generateParenthesis(n int) []string {
	ret := make([]string, 0)
	backtrackGP(&ret, "", 0, 0, n)
	return ret
}

func backtrackGP(ret *[]string, cur string, open int, close int, max int) {
	if len(cur) == max*2 {
		*ret = append(*ret, cur)
		return
	}
	if open < max {
		backtrackGP(ret, cur+"(", open+1, close, max)
	}
	if close < open {
		backtrackGP(ret, cur+")", open, close+1, max)
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
}
