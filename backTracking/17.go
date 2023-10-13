package backTracking

func letterCombinations(digits string) []string {
	ans := make([]string, 0)
	//用于临时记录每一个符合条件结果
	path := make([]byte, 0)
	if digits == "" {
		return ans
	}
	recode := map[int]string{
		2: "abc",
		3: "def",
		4: "ghi",
		5: "jkl",
		6: "mno",
		7: "pqrs",
		8: "tuv",
		9: "wxyz",
	}
	//字符串长度即深度
	var backtracking func(digits string, index int)
	backtracking = func(digits string, index int) {
		//终止条件是什么？
		if len(path) == len(digits) {
			ans = append(ans, string(path))
			return
		}

		//找到对应字符串
		digit := int(digits[index] - '0')
		str := recode[digit]
		for i := 0; i < len(str); i++ {
			//拼接结果
			path = append(path, str[i])
			//递归到下一层
			backtracking(digits, index+1)
			//回溯
			path = path[:len(path)-1]
		}
	}
	backtracking(digits, 0)
	return ans
}
