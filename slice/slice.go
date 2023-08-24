package slice

import "fmt"

func Expansion() {
	arr := make([]int, 0, 3)
	perCap := cap(arr)
	for i := 0; i < 100; i++ {
		arr = append(arr, i)
		curCap := cap(arr)
		if curCap > perCap {
			fmt.Printf("容量从：%d扩容到了：%d\n", perCap, curCap)
			fmt.Printf("arr:%v\n", arr)
			perCap = curCap
		}
	}
}
