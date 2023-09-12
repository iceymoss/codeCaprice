package main

import "fmt"

// 本题为考试单行多行输入输出规范示例，无需提交，不计分。
func main() {
	t := 0
	fmt.Scan(&t)
	ans := make([]string, 0)
	for i := 1; i <= t; i++ {
		a := 0
		b := 0
		fmt.Scan(&a, &b)
		if (a-1)%2 != 0 && (b-1)%2 == 0 {
			ans = append(ans, "Yes")
		} else if (a-1)%2 == 0 && (b-1)%2 != 0 {
			ans = append(ans, "Yes")
		} else {
			ans = append(ans, "No")
		}
	}
	for _, v := range ans {
		fmt.Println(v)
	}
}
