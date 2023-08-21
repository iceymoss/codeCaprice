package array

// func search(nums []int, target int) int {
//     //暴力
//     for i, v := range nums {
//         if v == target {
//             return i
//         }
//     }
//     return -1
// }

func search(nums []int, target int) int {
	//二分查找法
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
