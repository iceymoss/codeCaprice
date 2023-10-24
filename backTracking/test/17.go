package test

func letterCombinations(digits string) []string {
	ans := make([]string, 0)
	path := make([]byte, 0)
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
	var backtracking func(digits string, index int)
	backtracking = func(digits string, index int) {
		//终止条件
		if len(digits) == len(path) {
			ans = append(ans, string(path))
			return
		}

		if index > len(digits)-1 {
			return
		}
		digit := int(digits[index] - '0')
		str := recode[digit]
		for i := index; i < len(str); i++ {
			path = append(path, str[i])
			backtracking(digits, i+1)
			path = path[:len(path)-1]
		}

	}
	backtracking(digits, 0)
	return ans
}
