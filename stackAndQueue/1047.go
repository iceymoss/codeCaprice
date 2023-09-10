package stackAndQueue

func removeDuplicates(s string) string {
	//使用栈，先进后出，每次遍历时s时，拿到栈顶元素进行比较，相同则移除栈顶
	//不相同则加入栈，最后栈中就是去重相邻字符的结果
	bt := []byte(s)
	stack := make([]byte, 0)
	for i := 0; i < len(bt); i++ {
		if len(stack) == 0 {
			stack = append(stack, bt[i])
		} else {
			if stack[len(stack)-1] == bt[i] {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, bt[i])
			}
		}
	}
	return string(stack)
}

func removeDuplicatesV2(s string) string {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		// 栈不空 且 与栈顶元素不等
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			// 入栈
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
