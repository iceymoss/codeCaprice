package main

import (
	"fmt"
	"sort"
	"strconv"
)

// 将数组元素转为字符串进行排序
func maxCombination(nums []int) string {
	// 将数字转换为字符串数组
	strNums := make([]string, len(nums))
	for i, num := range nums {
		strNums[i] = strconv.Itoa(num)
	}

	// 使用自定义比较函数进行排序
	sort.Slice(strNums, func(i, j int) bool {
		return strNums[i]+strNums[j] > strNums[j]+strNums[i]
	})

	// 将排序后的字符串数组拼接在一起
	result := ""
	for _, strNum := range strNums {
		result += strNum
	}
	return result
}

// maxCombinationBackTracking 回溯法
func maxCombinationBackTracking(nums []int) string {
	ansMax := 0
	path := make([]int, 0)
	used := make([]bool, len(nums))
	var backTracking func()
	backTracking = func() {
		if len(path) == len(nums) {
			str := ""
			for _, v := range path {
				str += strconv.Itoa(v)
			}
			tmp, _ := strconv.Atoi(str)
			if tmp > ansMax {
				ansMax = tmp
			}
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backTracking()
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backTracking()
	return strconv.Itoa(ansMax)
}

func main() {
	nums := []int{1, 2, 30, 301, 9, 9, 101, 8}
	//fmt.Println(maxCombination(nums)) //3030121
	fmt.Println(maxCombinationBackTracking(nums))
}
