package main

import (
	"fmt"
	"strconv"
)

//Runtime: 4 ms, faster than 58.25% of Go online submissions for Count and Say.
//func countAndSay(n int) string {
//	var str = "1"
//	var getStr = func(str string) (next string) {
//		for i := 0; i < len(str); i++ {
//			if i == len(str)-1 {
//				next += "1" + string(str[i])
//			} else {
//				var n int = 0
//				var v = str[i]
//				for i < len(str) {
//					if str[i] == v {
//						n++
//						i++
//					} else {
//						next += strconv.Itoa(n) + string(v)
//						n = 0
//						i--
//						break
//					}
//				}
//				if n != 0 {
//					next += strconv.Itoa(n) + string(v)
//				}
//			}
//		}
//		return next
//	}
//	for i := 1; i < n; i++ {
//		str = getStr(str)
//	}
//	return str
//}

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Count and Say.
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	lastString := "1"
	for i := 2; i <= n; i++ {
		var results []byte
		lastByte := lastString[0]
		num := 1
		for j := range lastString {
			if j == 0 {
				continue
			}
			if lastString[j] == lastByte {
				num++
			} else {
				results = append(results, byte(num+48))
				results = append(results, lastByte)
				lastByte = lastString[j]
				num = 1
			}
		}
		results = append(results, byte(num+48))
		results = append(results, lastByte)
		lastString = string(results)
	}
	return lastString
}

func main() {
	fmt.Println(countAndSay(6))
}
