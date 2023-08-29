package towPointer

func removeElement(nums []int, val int) int {
	//暴力法，超时了
	ans := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			ans--
			for j := i + 1; j < len(nums); j++ {
				nums[j-1] = nums[j]
			}
			i--
		}
	}
	return ans
}

func removeElementV2(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
