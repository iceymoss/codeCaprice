package hashTable

import (
	"sort"
)

// 242. 有效的字母异位词
func IsAnagram(s string, t string) bool {
	//如果是字母异分词，字母的种类和出现次数一定是相同的，所以排序后的结果一定相同
	sr := []byte(s)
	tr := []byte(t)
	sort.Slice(sr, func(i, j int) bool {
		return sr[i] < sr[j]
	})

	sort.Slice(tr, func(i, j int) bool {
		return tr[i] < tr[j]
	})

	return string(sr) == string(tr)
}

func IsAnagramv2(s string, t string) bool {
	//使用map进行暴力: 时间：O(n) 空间：O(n)
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[int32]int)
	tMap := make(map[int32]int)
	for _, v := range s {
		sMap[v]++
	}
	for _, v := range t {
		tMap[v]++
	}

	for k, v := range sMap {
		if _, ok := tMap[k]; !ok || v != tMap[k] {
			return false
		}
	}
	return true
}

func IsAnagramv3(s string, t string) bool {
	//使用map,优化空间:时间：O(n),空间:O(1)
	//使用数组作为map,使用26个字母的槽位
	recode := [26]int{}
	for _, v := range s {
		recode[v-('a')]++
	}
	for _, v := range t {
		recode[v-('a')]--
	}
	return recode == [26]int{}
}
