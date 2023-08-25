package main

import (
	String "codeCaprice/string"
	"fmt"
)

func main() {
	arr := []byte("iceymoss")
	String.ReverseString(arr)
	fmt.Println(string(arr))
}
