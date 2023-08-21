package array

// minSubArrayLen 209. 长度最小的子数组
func minSubArrayLen(target int, nums []int) int {
	//暴力法
	//子数组不可大于len(nums)+1
	ans := len(nums) + 1
	for i := 0; i < len(nums); i++ {
		count := 0
		win := 0
		for j := i; j < len(nums); j++ {
			win += nums[j]
			count++
			if win >= target {
				ans = min(ans, count)
			}
		}
	}
	if ans == len(nums)+1 {
		return 0
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minSubArrayLenV2(target int, nums []int) int {
	//滑动窗口(双指针法)
	n := len(nums)
	l := 0
	ans := n + 1 //返回结果
	sum := 0
	//r表示右窗口, 向前推进
	for r := 0; r < len(nums); r++ {
		sum += nums[r]
		for sum >= target { //如果符合条件
			ans = min(ans, (r-l)+1)
			//l表示左窗口，当sum >= target时需要向右移动
			sum -= nums[l]
			l++
		}
	}
	if ans == n+1 {
		return 0
	}
	return ans
}
