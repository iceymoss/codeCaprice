package greedy

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if len(nums) < 2 {
		return n
	}
	ans := 1 //默认只有两个元素，如果相同ans = 1, 不相同ans = 2，解释：仅有一个元素或者含两个不等元素的序列也视作摆动序列。
	perDiff := nums[1] - nums[0]
	if perDiff != 0 {
		ans = 2
	}

	for i := 2; i < n; i++ {
		curDiff := nums[i] - nums[i-1]
		if perDiff > 0 && curDiff <= 0 || perDiff < 0 && curDiff >= 0 {
			ans++
			perDiff = curDiff
		}
	}
	return ans
}
