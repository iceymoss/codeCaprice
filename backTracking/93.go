package backTracking

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	ans := make([]string, 0)
	path := make([]string, 0)
	var backtracking func(s string, start int)
	backtracking = func(s string, start int) {
		if len(path) == 4 && start == len(s) {
			ip := strings.Join(path, ".")
			ans = append(ans, ip)
			return
		}

		if len(path) >= 4 && start < len(s) {
			return
		}

		for i := start; i < len(s); i++ {
			str := s[start : i+1]
			target, err := strconv.Atoi(str)
			if err != nil || len(str) != len(strconv.Itoa(target)) { //去除无效数字：如1.011.23.23 ->011
				return
			}

			if target > 255 || target < 0 {
				return
			}
			path = append(path, str)
			backtracking(s, i+1)
			path = path[:len(path)-1]
		}
	}
	backtracking(s, 0)
	return ans
}
