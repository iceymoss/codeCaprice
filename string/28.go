package string

func strStr(haystack, needle string) int {
	mainLen, subLen := len(haystack), len(needle)
	if mainLen < subLen {
		return -1
	}

	//滑动窗口，needle长度固定，然后我们要保持haystack[i:i+subLen]的窗口，并且从从左向右滑动
	//subLen为needle长度，所以我们i只能小于mainLen-subLen
	for i := 0; i < mainLen-subLen; i++ {
		if haystack[i:i+subLen] == needle {
			return i
		}
	}
	return -1
}

func findSubstringPosition(mainStr, subStr string) int {
	mainLen := len(mainStr)
	subLen := len(subStr)

	if mainLen < subLen {
		return -1
	}

	for i := 0; i <= mainLen-subLen; i++ {
		if mainStr[i:i+subLen] == subStr {
			return i
		}
	}

	return -1
}
