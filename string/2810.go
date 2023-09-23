package string

func finalString(s string) string {
	ans := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == 'i' {
			ReverseArr(ans)
			continue
		}
		ans = append(ans, s[i])
	}
	return string(ans)
}

func ReverseArr(arr []byte) {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}
