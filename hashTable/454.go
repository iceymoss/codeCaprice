package hashTable

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	//暴力法,超时的
	ans := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			for n := 0; n < len(nums3); n++ {
				for m := 0; m < len(nums4); m++ {
					if nums1[i]+nums2[j]+nums3[n]+nums4[m] == 0 {
						ans++
					}
				}
			}
		}
	}
	return ans
}

func fourSumCountV2(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	//使用map, key记录a+b的值，value记录a+b出现的参数
	ans := 0
	recode := make(map[int]int)
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			recode[v1+v2]++
		}
	}

	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			//a+b = -(c+d)
			ans += recode[-v3-v4]
		}
	}
	return ans
}
