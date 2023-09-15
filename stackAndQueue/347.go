package stackAndQueue

import (
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	//直接使用map不就行了吗
	ans := make([]int, 0)
	tag := make(map[int]int)
	for _, v := range nums {
		tag[v]++
	}

	for key, _ := range tag {
		ans = append(ans, key)
	}
	sort.Slice(ans, func(a, b int) bool {
		return tag[ans[a]] > tag[ans[b]]
	})
	return ans[:k]
}

//在这里，比较函数 func(a, b int) bool {...} 比较了 ans 中的两个元素 ans[a] 和 ans[b]。
//它的返回值指示了两个元素之间的顺序。具体来说，比较的规则是按照它们在 map_num 中出现的次数降序排列，也就是出现次数多的元素在前面。
//这样，ans 切片将按照 map_num 中元素出现的次数从高到低进行排序。
