package string

//459.重复的子字符串
/*
示例 1:

	输入: "abab"
	输出: True
	解释: 可由子字符串 "ab" 重复两次构成。
	示例 2:

	输入: "aba"
	输出: False
	示例 3:

	输入: "abcabcabcabc"
	输出: True
	解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)
#
*/

func RepeatedSubstringPattern(s string) bool {
	//暴力法:找到一个底层字符，拼接字符，然后对比
	tmp := ""
	for i := 0; i < len(s); i++ {
		tmp += string(s[i])
		tag := ""
		for len(tag) < len(s) {
			tag += tmp
		}
		if tag == s && tmp != s {
			return true
		}
	}
	return false
}

func RepeatedSubstringPatternV2(s string) bool {
	//暴力法:遍历字符串，找到子串:从左向右找子串。找到子串后然后向后遍历
	n := len(s)

	//i是一个偏移量
	for i := 1; i*2 < n; i++ {
		if n%i == 0 {
			macht := true
			for j := i; j < n; j++ {
				if s[j] != s[j-i] {
					macht = false
					break
				}
			}
			if macht {
				return true
			}
		}
	}
	return false
}
