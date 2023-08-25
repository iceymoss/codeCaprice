package string

//344. 反转字符串

func ReverseString(s []byte) {
	//双指针法，原地修改
	l, r := 0, len(s)-1
	for l <= r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
