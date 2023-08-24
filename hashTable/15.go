package hashTable

import (
	"fmt"
	"sort"
)

// 15. 三数之和
func threeSum(nums []int) [][]int {
	//a+b+c = 0  => map[a]， map[-(b+c)]中a = -(b+c)
	//map：key表示a的，value表示下标
	var result [][]int
	sort.Ints(nums)
	var ans [][]int //返回结果
	recode := make(map[int]int)
	for k := 0; k < len(nums); k++ {
		if nums[k] > 0 {
			return result
		}
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		recode[nums[k]] = k
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if index, ok := recode[-(nums[i] + nums[j])]; ok && index != i && index != j {
				arr := []int{nums[i], nums[j], nums[index]}
				sort.Ints(arr)
				ans = append(ans, arr)
			}
		}
	}

	visited := make(map[string]bool)
	//去重
	for _, row := range ans {
		rowStr := fmt.Sprintf("%s", row)
		if !visited[rowStr] {
			visited[rowStr] = true
			result = append(result, row)
		}
	}

	return result
}
