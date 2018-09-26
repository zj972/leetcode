package main

import "fmt"

//func longestPalindrome(s string) string {
//	var num string
//	for i := 0; i < len(s); i++ {
//		for j := i; j <= len(s); j++ {
//			if len(num) < palindroom(s[i:j]) {
//				num = s[i:j]
//			}
//		}
//	}
//	return num
//}
//
//func palindroom(s string) int {
//	for i := 0; i < len(s)/2; i++ {
//		if s[i] != s[len(s)-i-1] {
//			return -1
//		}
//	}
//	return len(s)
//}

func longestPalindrome(s string) string {
	str := []byte(s)
	if len(s) < 2 {
		return s
	}
	index := 0
	maxLen, start := 1, 0
	for index < len(s) {
		if maxLen/2 >= len(s)-index {
			break
		}
		j, k := index, index
		for k < len(s)-1 && str[k+1] == str[k] {
			k++
		}
		index = k + 1
		for k < len(s)-1 && j > 0 && str[j-1] == str[k+1] {
			j--
			k++
		}
		if k-j+1 > maxLen {
			start = j
			maxLen = k - j + 1
		}
	}
	return (string(str[start : start+maxLen]))
}

func main() {
	s := longestPalindrome("a")
	fmt.Println(s)
}
