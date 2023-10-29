package backTracking

func partition(s string) [][]string {
	ans := make([][]string, 0)
	path := make([]string, 0)
	var backtracking func(start int, s string)
	backtracking = func(start int, s string) {
		if start == len(s) { // 如果起始位置等于s的大小，说明已经找到了一组分割方案了
			ans = append(ans, append([]string{}, path...))
			return
		}

		for i := start; i < len(s); i++ {
			str := s[start : i+1]
			if isPalindrome(str) {
				path = append(path, str)
				backtracking(i+1, s)
				path = path[:len(path)-1]
			}
		}
	}
	backtracking(0, s)
	return ans
}

func isPalindrome(str string) bool {
	l, r := 0, len(str)-1
	for l < r {
		if str[l] != str[r] {
			return false
		}
		l++
		r--
	}
	return true
}
