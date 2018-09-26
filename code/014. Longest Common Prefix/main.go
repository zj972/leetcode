package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	var prefix string
	if len(strs) == 0 {
		prefix = ""
	} else {
		prefix = strs[0]
		for i := 1; i < len(strs); i++ {
			for strings.HasPrefix(strs[i], prefix) == false {
				prefix = string([]rune(prefix)[:len(prefix)-1])
			}
		}
	}
	return prefix
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"213", "213212", "21243223121"}))
}
