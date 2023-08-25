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

func threeSumV2(nums []int) [][]int {
	//先排序，为了更好的去重， 定一动二，双指针，也可以理解为三指针法
	sort.Ints(nums)
	var ans [][]int
	//a+b+c = 0
	//a = nums[i] b = nums[i+1] c = nums[len(nums)-1]
	for i := 0; i < len(nums)-2; i++ {
		if nums[0] > 0 {
			break
		}

		//nums[i]去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		n1 := nums[i]

		//双指针法
		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			if n1+n2+n3 == 0 {
				ans = append(ans, []int{n1, n2, n3})
				//l，r去重, 如果当前符合条件后，当前条件下：n1使用，n2使用，n3使用，所以如果n2或者n3不能再出现和上一层相同的值了
				for l < r && nums[l] == n2 {
					l++
				}
				for l < r && nums[r] == n3 {
					r--
				}

			} else if n1+n2+n3 > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return ans
}
