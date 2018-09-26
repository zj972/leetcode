package main

import "fmt"

func romanToInt(s string) int {
	x := 0
	num := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		switch string(s[i]) {
		case "I":
			num[i] = 1
		case "V":
			num[i] = 5
		case "X":
			num[i] = 10
		case "L":
			num[i] = 50
		case "C":
			num[i] = 100
		case "D":
			num[i] = 500
		case "M":
			num[i] = 1000
		}
	}
	for j, rules := 0, 0; j < len(num); j++ {
		if j == len(num)-1 {
			if rules != 0 {
				x += (rules + 1) * num[j]
				rules = 0
			} else {
				x += num[j]
			}
		} else if num[j] > num[j+1] {
			if rules != 0 {
				x += (rules + 1) * num[j]
				rules = 0
			} else {
				x += num[j]
			}
		} else if num[j] == num[j+1] {
			rules++
		} else {
			if rules != 0 {
				x -= (rules + 1) * num[j]
				rules = 0
			} else {
				x -= num[j]
			}
		}
	}
	return x
}

func main() {
	fmt.Println(romanToInt("MMCCCXXIII"))
}
