package hashTable

// 383. 赎金信
func canConstruct(ransomNote string, magazine string) bool {
	//map：key表示magazine对应字符，value表示magazine对应字符出现的次数
	recode := make(map[int32]int)
	for _, v := range magazine {
		recode[v]++
	}

	for _, v := range ransomNote {
		if recode[v] == 0 {
			return false
		}
		if _, ok := recode[v]; ok && recode[v] > 0 {
			recode[v]--
		}
	}
	return true
}

// 优化空间
func canConstructV2(ransomNote string, magazine string) bool {
	recode := make([]int, 26)
	for _, v := range magazine {
		recode[v-'a']++
	}
	for _, v := range ransomNote {
		recode[v-'a']--
		if recode[v-'a'] < 0 {
			return false
		}
	}
	return true
}
