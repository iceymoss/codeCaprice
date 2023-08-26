package string

// 请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
func replaceSpace(s string) string {
	tag := "%20"
	var ans []byte
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			ans = append(ans, []byte(tag)...)
		} else {
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}

// 时间换空间，原地修改
// 原地修改
func replaceSpacV2(s string) string {
	b := []byte(s)
	length := len(b)
	spaceCount := 0
	// 计算空格数量
	for _, v := range b {
		if v == ' ' {
			spaceCount++
		}
	}
	// 扩展原有切片
	resizeCount := spaceCount * 2
	tmp := make([]byte, resizeCount)
	b = append(b, tmp...)
	i := length - 1
	j := len(b) - 1
	for i >= 0 {
		if b[i] != ' ' {
			b[j] = b[i]
			i--
			j--
		} else {
			b[j] = '0'
			b[j-1] = '2'
			b[j-2] = '%'
			i--
			j = j - 3
		}
	}
	return string(b)
}
