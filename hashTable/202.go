package hashTable

// 202. 快乐数
func isHappy(n int) bool {
	//为了防止，无限循环，需要使用map来记录一下
	if n < 0 {
		return false
	}
	recode := make(map[int]bool)
	for n != 1 && !recode[n] {
		recode[n] = true
		n = step(n)
	}
	return n == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}
