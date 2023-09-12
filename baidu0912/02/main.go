package main

import "fmt"

// 本题为考试单行多行输入输出规范示例，无需提交，不计分。
func main() {
	t := 0
	fmt.Scan(&t)
	nice := make([]int, 0)
	bad := make([]int, 0)
	for i := 1; i <= t; i++ {
		n := 0
		fmt.Scan(&n)
		nice = append(nice, n)
	}
	for i := 1; i <= t; i++ {
		n := 0
		fmt.Scan(&n)
		bad = append(bad, n)
	}

	diff := make([]int, t)
	for i := 0; i < t; i++ {
		diff[i] = nice[i] - bad[i]
	}
	sum1, sum2 := 0, 0
	for _, v := range diff {
		if v < 0 {
			sum1 += v
		} else {
			sum2 += v
		}
	}

	if -1*sum1 > sum2 {
		fmt.Println(sum1 * -1)
	} else {
		fmt.Println(sum2)
	}
}
