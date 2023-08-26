package string

import (
	"fmt"
)

/*
给定一个字符串，逐个翻转字符串中的每个单词。

示例 1：
输入: "the sky is blue"
输出: "blue is sky the"

示例 2：
输入: "  hello world!  "
输出: "world! hello"
解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。

示例 3：
输入: "a good   example"
输出: "example good a"
解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
*/

func ReverseWords(s string) string {
	bt := []byte(s)

	//1.去除冗余空格：移除前面、中间、后面存在的多余空格
	slow := 0
	for i := 0; i < len(bt); i++ {
		//去除冗余空格，i找到不为空格的字符，slow用于记录bt[i]移动的位置
		if bt[i] != ' ' {
			//保留单词直接的空格
			if slow != 0 {
				bt[slow] = ' '
				slow++
			}
			//将多余的空格去掉：将不为空格的字符向前移动
			for i < len(bt) && bt[i] != ' ' {
				bt[slow] = bt[i]
				i++
				slow++
			}
		}
	}
	bt = bt[0:slow]

	//2.反转整个字符串
	Reverse(bt)

	//3.反转单词
	start := 0
	for i := 0; i <= len(bt); i++ {
		if i == len(bt) || bt[i] == ' ' {
			Reverse(bt[start:i])
			start = i + 1
		}
	}
	return string(bt)
}

func Reverse(bt []byte) {
	l, r := 0, len(bt)-1
	for l < r {
		bt[l], bt[r] = bt[r], bt[l]
		l++
		r--
	}
}

func OutEmptyFormat(str *[]byte) {
	//快慢指针，去除冗余字符
	//slow用来记录去除冗余后的字符
	//fast用于判断冗余空格
	slow, fast := 0, 0

	//去除头部冗余：将fast指向不为空格的地方
	for fast < len(*str) && (*str)[fast] == ' ' {
		fast++
	}

	//去除单词间冗余
	for ; fast < len(*str); fast++ {
		if fast-1 > 0 && (*str)[fast-1] == (*str)[fast] && (*str)[fast] == ' ' {
			continue
		}
		//当fast遇到不为空的字符后，需要将
		(*str)[slow] = (*str)[fast]
		slow++
	}

	//去除尾部冗余
	if slow-1 > 0 && (*str)[slow-1] == ' ' {
		*str = (*str)[:slow-1]
	} else {
		*str = (*str)[:slow]
	}
	fmt.Println(string(*str))
}
