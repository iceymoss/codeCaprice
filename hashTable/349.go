package hashTable

// 349. 两个数组的交集
func intersection(nums1 []int, nums2 []int) (intersection []int) {
	//哈希表:时间：n 空间：n
	recode := make(map[int]bool)
	for _, v := range nums2 {
		recode[v] = true
	}
	for _, v := range nums1 {
		if recode[v] {
			intersection = append(intersection, v)
			recode[v] = false
		}
	}
	return
}

// 空间优化，使用struct{}{}几乎不占内存
func intersectionV2(nums1 []int, nums2 []int) (intersection []int) {
	recode := make(map[int]struct{})
	for _, v := range nums1 {
		recode[v] = struct{}{}
	}
	for _, v := range nums2 {
		if _, ok := recode[v]; ok {
			intersection = append(intersection, v)
			delete(recode, v)
		}
	}
	return
}
