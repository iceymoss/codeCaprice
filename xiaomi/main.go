package main

import "fmt"

func getXiaoMiCount(str string) int {
	bt := []byte(str)
	tag := "xiaomi"
	ans := 1
	path := make([]byte, 0)
	var backtraking func(index int)
	backtraking = func(index int) {
		//递归终止条件
		if string(path) == tag {
			ans++
			fmt.Println(string(path))
			return
		}

		if len(path) == len(tag) {
			fmt.Println(string(path))
			return
		}

		for i := index; i < len(bt); i++ {
			path = append(path, bt[i])
			backtraking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtraking(0)
	return ans
}

func getXiaoMiCountV2(str string) int {
	bt := []byte(str)
	tag := make(map[byte]bool)
	for _, v := range []byte(str) {
		tag[v] = true
	}
	queue := make([]byte, 0)

	for _, b := range bt {
		if tag[b] {
			queue = append(queue, b)
		}
	}
	fmt.Println(string(queue))
	return 0
}

func main() {
	var b []byte
	fmt.Scan(&b)

	getXiaoMiCountV2(string(b))
	//res := getXiaoMiCount(string(b))
	//fmt.Println(res)

}
