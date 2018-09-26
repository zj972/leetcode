package main

import (
	"fmt"
	"strings"
)

//Runtime: 0 ms
//Your runtime beats 100.00 % of golang submissions.
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	digitMap := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	add := func(list []string, str string) []string {
		outList := []string{}
		for n := 0; n < len(digitMap[str]); n++ {
			midList := []string{}
			for i := 0; i < len(list); i++ {
				midList = append(midList, list[i]+string(digitMap[str][n]))
			}
			outList = append(outList, midList...)
		}
		return outList
	}
	output := []string{""}
	digitList := strings.Split(digits, "")
	for _, val := range digitList {
		output = add(output, val)
	}
	return output
}

func main() {
	fmt.Println(letterCombinations("23"))
}
