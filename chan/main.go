package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//fmt.Println([]int{})
	//
	//s1 := new([]int)
	//fmt.Println(s1)
	//
	//s2 := make([]int, 0)
	//fmt.Println(s2)
	n1 := node1{"iceymoss", 1}
	n2 := node1{"iceymoss", 1}
	fmt.Println(n1 == n2)
	s := struct{}{}
	fmt.Println(unsafe.Sizeof(s))
}

type node1 struct {
	name string
	age  int
}

type node2 struct {
	name string
	age  int
}

//type node1 struct {
//	arr []int
//	m map[string]string
//}
//
//type node2 struct {
//	arr []int
//	m map[string]string
//}
