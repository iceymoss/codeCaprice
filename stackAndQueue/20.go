package stackAndQueue

// ({})[]
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	m := make(map[byte]byte)
	m[')'] = '('
	m[']'] = '['
	m['}'] = '{'

	bt := []byte(s)
	stack := make([]byte, 0)
	for i := 0; i < len(bt); i++ {
		if _, ok := m[bt[i]]; !ok {
			stack = append(stack, bt[i])
		} else {
			if len(stack) > 0 {
				if m[bt[i]] != stack[len(stack)-1] {
					return false
				}
				stack = stack[:len(stack)-1]
			} else { //右括号多了的情况直接返回
				return false
			}
		}
	}
	return len(stack) == 0
}
