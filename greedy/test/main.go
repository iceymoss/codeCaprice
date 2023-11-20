package main

import (
	"codeCaprice/greedy"
	"fmt"
)

func main() {
	//fmt.Println(greedy.FindContentChildren([]int{1, 2, 3}, []int{1, 1}))
	fmt.Println(greedy.MaxSubArrayV2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
