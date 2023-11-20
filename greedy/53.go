package greedy

import "math"

func MaxSubArray(nums []int) int {
	ans := math.MinInt
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := i; j < len(nums); j++ {
			count += nums[j]
			if count > ans {
				ans = count
			}
		}
	}
	return ans
}

// 贪心法：局部最优到全局最优，遍历数组，局部最优：当前“连续和”为负数的时候立刻放弃，
//从下一个元素重新计算“连续和”，因为负数加上下一个元素 “连续和”只会越来越小。
//例如：x为负数，y为正数， 一定有 x+y < y

func MaxSubArrayV2(nums []int) int {
	ans := math.MinInt
	count := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if ans < count {
			ans = count
		}
		if count < 0 {
			count = 0
		}
	}
	return ans
}
