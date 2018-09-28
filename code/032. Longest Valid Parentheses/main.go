package main

import (
	"fmt"
)

//Runtime: 2000 ms, faster than 0.00% of Go online submissions for Longest Valid Parentheses.
//func longestValidParentheses(s string) int {
//	list := []rune(s)
//	var max int
//	for i := 0; i < len(list); i++ {
//		var n int
//		for j := i; j < len(list); j++ {
//			switch string(list[j]) {
//			case "(":
//				n++
//			case ")":
//				n--
//			}
//			if n < 0 {
//				break
//			} else if n == 0 {
//				if j-i+1 > max {
//					max = j - i + 1
//				}
//			}
//		}
//	}
//	return max
//}

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Valid Parentheses.
func longestValidParentheses(s string) int {
	stack := []int{-1}
	temp := 0
	longest := 0
	for pos, char := range s {
		if char == '(' {
			stack = append(stack, pos)
		} else {
			stack = stack[0 : len(stack)-1]
			if len(stack) > 0 {
				temp = pos - stack[len(stack)-1]
				if temp > longest {
					longest = temp
				}
			} else {
				stack = append(stack, pos)
			}
		}
	}
	return longest
}

func main() {
	var data = "(()(("
	fmt.Println(longestValidParentheses(data))
}
