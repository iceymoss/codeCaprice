package string

// 剑指 Offer 58 - II. 左旋转字符串
func reverseLeftWords(s string, n int) string {
	ans := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if i == n {
			ans = append([]byte(s[i:]), []byte(s[0:i])...)
		}
	}
	return string(ans)
}

func reverseLeftWordsV2(s string, n int) string {
	bt := []byte(s)
	for i := 0; i < len(s); i++ {
		if i == n {
			bt = append(bt, bt[0:i]...)
			bt = bt[i:]
		}
	}
	return string(bt)
}

func reverseLeftWordsV3(s string, n int) string {
	bt := []byte(s)
	l := len(bt)

	tmp := ""
	for i := 0; i < len(s); i++ {
		if i == n {
			tmp = string(bt[:i])
		}
	}
	lenght := len(tmp)
	for j := n; j < len(bt); j++ {
		bt[j-lenght] = bt[j]
	}

	index := l - n
	for i, v := range tmp {
		bt[index+i] = byte(v)
	}
	return string(bt)
}

/*
这道题目也非常类似，依然可以通过局部反转+整体反转 达到左旋转的目的。

具体步骤为：

反转区间为前n的子串
反转区间为n到末尾的子串
反转整个字符串
最后就可以达到左旋n的目的，而不用定义新的字符串，完全在本串上操作。
*/
func reverseLeftWordsV4(s string, n int) string {
	bt := []byte(s)
	revereBytes(bt[:n])
	revereBytes(bt[n:])
	revereBytes(bt)
	return string(bt)
}

func revereBytes(b []byte) {
	l, r := 0, len(b)-1
	for l < r {
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
}
