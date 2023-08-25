package hashTable

import "sort"

func fourSum(nums []int, target int) [][]int {
	//思路：先排序，方便去重，定一，次定一，移二(双指针）
	var ans [][]int
	//a = nums[i], b = nums[j], c = nums[l], d = nums[r]
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ { //注意数组越界
		if target <= 0 && nums[0] > 0 {
			break
		}
		//nums[i]去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		n1 := nums[i]
		for j := i + 1; j < len(nums)-2; j++ { //注意数组越界
			n2 := nums[j]
			//nums[j]去重
			if j > i+1 && n2 == nums[j-1] {
				continue
			}
			l, r := j+1, len(nums)-1 //双指针
			for l < r {
				n3, n4 := nums[l], nums[r]
				if n1+n2+n3+n4 == target {
					ans = append(ans, []int{n1, n2, n3, n4})
					for l < r && n3 == nums[l] {
						l++
					}
					for l < r && n4 == nums[r] {
						r--
					}
				} else if n1+n2+n3+n4 > target {
					r--
				} else {
					l++
				}
			}
		}
	}
	return ans
}
