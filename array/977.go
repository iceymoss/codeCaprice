package array

//977.有序数组的平方

// func sortedSquares(nums []int) []int {
//     //暴力法
//     // ans := make([]int, len(nums))
//     for k, v := range nums {
//         nums[k] = v*v
//     }
//     sort.Ints(nums)
//     return nums
// }

func sortedSquares(nums []int) []int {
	//双指针法，对于已经排序好的数组，平方值一定在数组的两端（不管是正数还是负数)
	ans := make([]int, len(nums))
	//i从0开始，j从len(nums)-1开始
	i, j, k := 0, len(nums)-1, len(nums)-1
	for i <= j {
		ni, nj := nums[i]*nums[i], nums[j]*nums[j]
		if ni < nj {
			ans[k] = nj
			j--
		} else {
			ans[k] = ni
			i++
		}
		k--
	}
	return ans
}
