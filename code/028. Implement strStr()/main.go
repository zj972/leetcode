package main

import (
	"fmt"
)

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Implement strStr().
//func strStr(haystack string, needle string) int {
//	return strings.Index(haystack, needle)
//}

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Implement strStr().
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	s := []rune(haystack)
	n := []rune(needle)
	for i := 0; i <= len(s)-len(n); i++ {
		if s[i] == n[0] {
			j := 0
			for ; j < len(n); j++ {
				if n[j] != s[i+j] {
					break
				}
			}
			if j == len(n) {
				return i
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("hello", "ll"))
}
