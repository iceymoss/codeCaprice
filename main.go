package main

import "fmt"

//func foo(arg_val int) *int {
//
//	//var foo_val1 int = 11
//	//var foo_val2 int = 12
//	//var foo_val3 int = 13
//	//var foo_val4 int = 14
//	//var foo_val5 int = 15
//	//
//	////此处循环是防止go编译器将foo优化成inline(内联函数)
//	////如果是内联函数，main调用foo将是原地展开，所以foo_val1-5相当于main作用域的变量
//	////即使foo_val3发生逃逸，地址与其他也是连续的
//	//for i := 0; i < 5; i++ {
//	//	println(&arg_val, &foo_val1, &foo_val2, &foo_val3, &foo_val4, &foo_val5)
//	//}
//	//
//	////返回foo_val3给main函数
//	//return &foo_val3
//
//	var foo_val1 *int = new(int)
//	var foo_val2 *int = new(int)
//	var foo_val3 *int = new(int)
//	var foo_val4 *int = new(int)
//	var foo_val5 *int = new(int)
//
//	//此处循环是防止go编译器将foo优化成inline(内联函数)
//	//如果是内联函数，main调用foo将是原地展开，所以foo_val1-5相当于main作用域的变量
//	//即使foo_val3发生逃逸，地址与其他也是连续的
//	for i := 0; i < 5; i++ {
//		println(arg_val, foo_val1, foo_val2, foo_val3, foo_val4, foo_val5)
//	}
//
//	//返回foo_val3给main函数
//	return foo_val3
//}

func main() {
	//main_val := foo(666)
	//
	//println(*main_val, main_val)

	//逃逸分析
	//data4 := []int{1, 3}
	//data4[0] = 2

	//data := []interface{}{100, 200}
	//data[0] = 100
	//
	//data1 := make(map[string]interface{})
	//data1["key"] = 200
	//
	//data2 := make(map[interface{}]interface{})
	//data2[100] = 200
	//
	//data3 := make(map[string][]string)
	//data3["key"] = []string{"value"}
	//
	//a := 10
	//data4 := []*int{nil}
	//data4[0] = &a
	//
	//data5 := 10
	//f := foo
	//f(&data5)

	s := "1234"
	for _, v := range s {
		fmt.Println(v - 48)
	}

}

func foo(a *int) {
	return
}
