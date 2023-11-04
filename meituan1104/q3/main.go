package main

import "fmt"

func main() {
	var n, q int
	fmt.Scan(&n, &q)

	var str string
	fmt.Scan(&str)

	inds := make([]int, q)
	for i := 0; i < q; i++ {
		var index int
		fmt.Scan(&index)
		inds[i] = index
	}

	t := str
	bt := []byte(str)

	for len(bt) > 0 {
		if len(bt) != 0 {
			bt = bt[:len(bt)-1]
			t += string(bt)
		}
		if len(bt) != 0 {
			bt = bt[1:]
			t += string(bt)
		}
	}
	fmt.Println("结果：", t)
	for _, v := range inds {
		fmt.Println(string(t[v-1]))
	}
}
