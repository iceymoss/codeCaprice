package backTracking

import (
	"sort"
	"strconv"
)

//思路：必须排序，记录所有的组合，然后使用map去重

func SubsetsWithDup(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	used := make(map[string]bool)
	str := ""
	sort.Ints(nums)
	var backtracking func(index int, str *string)
	backtracking = func(index int, str *string) {
		if index <= len(nums) {
			used[*str] = true
			ans = append(ans, append([]int{}, path...))
		}

		if index > len(nums) {
			return
		}

		for i := index; i < len(nums); i++ {
			s := strconv.Itoa(nums[i])
			length := len(s)
			*str += s
			if !used[*str] {
				path = append(path, nums[i])
				backtracking(i+1, str)
				path = path[:len(path)-1]
			}
			*str = (*str)[:len(*str)-length]
		}
	}
	backtracking(0, &str)
	return ans
}

//优化后：
//思路：先对数组进行排序方便去重，排序后画图出来可以知道[1,1,2,2]中，在递归到nums[0]中
//递归的nums[1]的结果肯定会出现在nums[0]的结果集中，所以同层之间的数据不能相同

func SubsetsWithDupV1(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	sort.Ints(nums)
	var backtracking func(index int)
	backtracking = func(index int) {
		if index <= len(nums) {
			ans = append(ans, append([]int{}, path...))
		}

		if index > len(nums) {
			return
		}

		for i := index; i < len(nums); i++ {
			if i != index && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(0)
	return ans
}
