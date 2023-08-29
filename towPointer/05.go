package towPointer

//剑指 Offer 05. 替换空格

func replaceSpace(s string) string {
	//暴力法
	bt := []byte(s)
	ans := make([]byte, 0)
	for i := 0; i < len(bt); i++ {
		if bt[i] == ' ' {
			ans = append(ans, []byte("%20")...)
		} else {
			ans = append(ans, bt[i])
		}
	}
	return string(ans)
}
