/*
1.关于数字的位置关系以及相关算法可以用位的思想来考虑
*/
package main

import ()

//func isPalindrome(x int) bool {
//	if x < 0 {
//		return false
//	} else if x == 0 {
//		return true
//	} else {
//		str := strings.Split(strconv.Itoa(x), "")
//		for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
//			str[i], str[j] = str[j], str[i]
//		}
//		s := strings.Join(str, "")
//		if y, _ := strconv.Atoi(s); y == x {
//			return true
//		} else {
//			return false
//		}
//	}
//}

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	rev := 0
	for x > rev {
		rev = rev*10 + x%10
		x /= 10
	}
	return x == rev || x == rev/10
}

func main() {
	isPalindrome(324)
}
