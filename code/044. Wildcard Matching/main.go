package main

import (
	"fmt"
)

//Time Limit Exceeded
//func isMatch(s string, p string) bool {
//	if s == "" && p == "" {
//		return true
//	}
//	if p == "" {
//		return false
//	}
//	if p[0] == 42 {
//		if len(p) == 1 {
//			return true
//		}
//		if s == "" {
//			return isMatch(s, p[1:])
//		} else {
//			for i := 0; i < len(s); i++ {
//				if isMatch(s[i:], p[1:]) {
//					return true
//				}
//			}
//			return false
//		}
//	} else if p[0] == 63 {
//		if s == "" {
//			return false
//		}
//		return isMatch(s[1:], p[1:])
//	} else {
//		if s == "" {
//			return false
//		}
//		if s[0] == p[0] {
//			return isMatch(s[1:], p[1:])
//		} else {
//			return false
//		}
//	}
//	return false
//}

//Runtime: 12 ms, faster than 97.44% of Go online submissions for Wildcard Matching.
func isMatch(s string, p string) bool {
	slen, plen := len(s), len(p)
	scur, pcur, sstar, pstar := 0, 0, -1, -1

	for scur < slen {
		if pcur < plen && (s[scur] == p[pcur] || p[pcur] == '?') {
			scur++
			pcur++
		} else if pcur < plen && p[pcur] == '*' {
			pstar = pcur
			pcur++
			sstar = scur
		} else if pstar >= 0 {
			pcur = pstar + 1
			sstar++
			scur = sstar
		} else {
			return false
		}
	}

	for pcur < plen && p[pcur] == '*' {
		pcur++
	}

	return pcur == plen
}

//Runtime: 8 ms, faster than 100.00% of Go online submissions for Wildcard Matching.
//func isMatch(s string, p string) bool {
//	slen, plen := len(s), len(p)
//	i, j := 0, 0
//	iStar, jStar := -1, -1
//
//	for ; i < slen; i, j = i+1, j+1 {
//		if j < plen && p[j] == '*' {
//			iStar = i
//			jStar = j
//			i--
//			continue
//		}
//		if j == plen || (s[i] != p[j] && p[j] != '?') {
//			if iStar >= 0 {
//				i = iStar
//				iStar++
//				j = jStar
//			} else {
//				return false
//			}
//		}
//	}
//	for ; j < plen && p[j] == '*'; j++ {
//	}
//
//	return j == plen
//}

func main() {
	fmt.Println(isMatch("aa", "a"))
}
