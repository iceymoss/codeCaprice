package main

import "fmt"

func toggle(char byte) byte {
	if char >= 'A' && char <= 'Z' {
		return char + 'a' - 'A'
	} else if char >= 'a' && char <= 'z' {
		return char - 'a' + 'A'
	}
	return char
}

func main() {
	var nums []string
	for {
		var num string
		_, err := fmt.Scanf("%s", &num)
		if err != nil {
			break
		}
		nums = append(nums, num)
	}
	for _, v := range nums {
		fmt.Println(string(toggle([]byte(v)[0])))
	}

}
