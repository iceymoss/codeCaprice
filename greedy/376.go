package greedy

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if len(nums) < 2 {
		return n
	}
	ans := 1
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
