package string

func reverseStr(s string, k int) string {
	bt := []byte(s)
	lenght := len(bt)
	for i := 0; i < len(bt); i += 2 * k {
		// 1. 每隔 2k 个字符的前 k 个字符进行反转
		// 2. 剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符
		if i+k < lenght {
			change(bt[i : i+k])
		} else {
			change(bt[i:lenght])
		}
	}
	return string(bt)

}

func change(bt []byte) {
	l, r := 0, len(bt)-1
	for l < r {
		bt[l], bt[r] = bt[r], bt[l]
		l++
		r--
	}
}
