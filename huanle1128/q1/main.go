package main

import "fmt"

func BubbleSort(arr []int, sortRule bool) []int {
	for i := 0; i < len(arr); i++ {
		tag := false
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] && sortRule || arr[j] < arr[j+1] && !sortRule {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				tag = true
			}
		}
		if !tag {
			break
		}
	}
	return arr
}

func main() {
	nums := []int{2, 3, 6, 1, 96, 13, 54}
	fmt.Println(BubbleSort(nums, false))

}
