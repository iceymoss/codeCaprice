package hashTable

func twoSum(nums []int, target int) []int {
	//暴力法
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSumV2(nums []int, target int) []int {
	//使用map
	m := make(map[int]int) //
	for i := 0; i < len(nums); i++ {
		if key, ok := m[target-nums[i]]; ok {
			return []int{key, i}
		}
		m[target-nums[i]] = i
	}
	return nil
}

//
