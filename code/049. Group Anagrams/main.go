package main

import (
	"fmt"
	"sort"
	"strings"
)

//拆分为字符串数组排序
//Runtime: 340 ms, faster than 46.46% of Go online submissions for Group Anagrams.
//func groupAnagrams(strs []string) [][]string {
//	tmp := make(map[string]int)
//	ret := [][]string{}
//	for i := 0; i < len(strs); i++ {
//		t := strSort(strs[i])
//		if tmp[t] != 0 {
//			ret[tmp[t]-1] = append(ret[tmp[t]-1], strs[i])
//		} else {
//			tmp[t] = len(ret) + 1
//			ret = append(ret, []string{strs[i]})
//		}
//	}
//	return ret
//}
//
//func strSort(str string) string {
//	tmp := strings.Split(str, "")
//	sort.Strings(tmp)
//	return strings.Join(tmp, "")
//}

//自己实现字符串排序
//Runtime: 292 ms, faster than 95.28% of Go online submissions for Group Anagrams.
//type SortByte []byte
//
//func (a SortByte) Len() int           { return len(a) }
//func (a SortByte) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
//func (a SortByte) Less(i, j int) bool { return a[i] < a[j] }
//
//func groupAnagrams(strs []string) (r [][]string) {
//	km := make(map[string]int)
//	i := 0
//	for _, v := range strs {
//		va := []byte(v)
//		sort.Sort(SortByte(va))
//		k := string(va)
//		if kl, ok := km[k]; !ok {
//			km[k] = i
//			r = append(r, []string{v})
//			i++
//		} else {
//			r[kl] = append(r[kl], v)
//		}
//	}
//	return r
//}

//通过质数求唯一值
//Runtime: 264 ms, faster than 99.21% of Go online submissions for Group Anagrams.
func groupAnagrams(strs []string) [][]string {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103}
	f := map[int][]string{}
	a := []int{}
	for _, v := range strs {
		sum := 1
		for i := 0; i < len(v); i++ {
			sum *= primes[int(v[i])-'a']
		}
		if f[sum] == nil {
			f[sum] = []string{v}
			a = append(a, sum)
		} else {
			f[sum] = append(f[sum], v)
		}
	}
	res := make([][]string, len(a))
	for i, sa := range a {
		res[i] = f[sa]
	}
	return res
}

func main() {
	var data = []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(data))
}
