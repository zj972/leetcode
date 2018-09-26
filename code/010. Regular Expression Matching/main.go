package main

import (
	"fmt"
)

//func isMatch(s string, p string) bool {
//  if p==s{
//    return true
//  }
//
//  b,_:= regexp.MatchString("^"+p+"$",s)
//  return b
//}

//func isMatch(s string, p string) bool {
//  if len(p) == 0 {
//    return len(s) == 0
//  }
//
//  if len(p) > 1 && p[1] == '*' {
//    return isMatch(s, p[2:]) || (len(s)>0 && (s[0] == p[0] || p[0]=='.') && isMatch(s[1:], p))
//  } else {
//    return (len(s) > 0 && (s[0] == p[0] || p[0] == '.') && isMatch(s[1:], p[1:]))
//  }
//}

func isMatch(s string, p string) bool {
	n, m := len(s), len(p)
	M := make([]bool, n+1)

	var star bool // is there a succeeding '*' after p[j]
	//var prev bool // previous value of M[i+1]
	M[n] = true // empty chars at the end of s & p
	for j := m - 1; j > -1; j-- {
		if p[j] == '*' {
			star = true
			continue
		}
		if !star {
			//match only one char of s with p[j]
			for i := 0; i < n; i++ {
				M[i] = (p[j] == '.' || s[i] == p[j]) && M[i+1]
			}
			//if M[0]{
			//	return false
			//}
			//for i := n; i > 0; i-- {
			//	if M[i] {
			//		if i == 1 && !M[0] {
			//			M[i - 1] = (p[j] == '.' || s[i - 1] == p[j]) && M[i]
			//		} else if i != 1 {
			//			if M[i - 1] {
			//				continue
			//			}else {
			//				M[i - 1] = (p[j] == '.' || s[i - 1] == p[j]) && M[i]
			//				break
			//			}
			//		} else {
			//			break
			//		}
			//	}
			//}
			M[n] = false // empty char does not match with p[j]
		} else {
			// match 0 or more chars s with p[j]
			for i := n - 1; i > -1; i-- {
				M[i] = M[i] || ((p[j] == '.' || s[i] == p[j]) && M[i+1])
			}
		}
		star = false
		//fmt.Println(M)
	}
	return M[0]
}

func main() {
	fmt.Println(isMatch("abcde", ".b*c*a*d."))
}
