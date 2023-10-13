package main

import (
	"fmt"
)

func countXiaomiSubsequences(s string, target string) int {
	// 创建一个二维切片用于存储子问题的解决方案
	dp := make([][]int, len(target)+1)
	for i := range dp {
		dp[i] = make([]int, len(s)+1)
	}

	// 初始化第一行，表示从s中找到空字符串的方式只有1种，即什么都不选
	for i := 0; i <= len(s); i++ {
		dp[0][i] = 1
	}

	// 填充二维表格
	for i := 1; i <= len(target); i++ {
		for j := 1; j <= len(s); j++ {
			// 如果s[j-1]和target[i-1]匹配，可以选择s[j-1]也可以不选
			if s[j-1] == target[i-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				// 否则只能选择不选s[j-1]
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[len(target)][len(s)]
}

func main() {
	s := "I love xiaomi, i often visit mi.com to buy phone"
	target := "xiaomi"
	result := countXiaomiSubsequences(s, target)
	fmt.Println(result)
}
