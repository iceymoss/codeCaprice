package greedy

import "sort"

func FindContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	index := 0
	for i := 0; i < len(s); i++ {
		if index < len(g) && g[index] <= s[i] {
			index++
		}
	}
	return index
}
