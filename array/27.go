package array

func removeElement(nums []int, val int) int {
	//暴力法
	ans := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val { //集体数组向前移动一位
			ans++
			for j := i + 1; j < len(nums); j++ {
				nums[j-1] = nums[j]
			}
			i-- // 因为下标i以后的数值都向前移动了一位，所以i也向前移动一位
		}
	}
	return ans
}

// func removeElement(nums []int, val int) int {
//     //快慢指针
//     slow := 0
//     for fast := 0; fast < len(nums); fast++ {
//         if nums[fast] != val {
//             nums[slow] = nums[fast]
//             slow++
//         }
//     }
//     return slow
// }
