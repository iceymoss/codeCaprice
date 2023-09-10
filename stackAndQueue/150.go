package stackAndQueue

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	if len(tokens) == 1 {
		return stringToInt(tokens[0])
	}
	stack := make([]string, 0)
	ans := 0
	for i := 0; i < len(tokens); i++ {
		//一定是运算符
		if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
			if len(stack) >= 2 {
				topNoe := stringToInt(stack[len(stack)-1])
				topTow := stringToInt(stack[len(stack)-2])
				stack = stack[:len(stack)-2]
				total := 0
				if tokens[i] == "+" {
					total = topTow + topNoe
				} else if tokens[i] == "-" {
					total = topTow - topNoe
				} else if tokens[i] == "*" {
					total = topNoe * topTow
				} else if tokens[i] == "/" {
					total = topTow / topNoe
				}
				ans = total
				stack = append(stack, strconv.Itoa(total))
			}
		} else { //如果是数字则直接入栈
			stack = append(stack, tokens[i])
		}
	}
	return ans
}

// func stringToInt(s string) int {
// 	tag, _ := strconv.Atoi(s)
// 	return tag
// }

// 输入：tokens = ["2","1","+","3","*"]
// 输出：9
// 解释：该算式转化为常见的中缀算术表达式为：((2 + 1) * 3) = 9
func evalRPNV2(tokens []string) int {
	//使用栈，根据运算符计算栈顶两个元素，然后将值再放回栈顶
	if len(tokens) == 1 {
		return stringToInt(tokens[0])
	}
	stack := make([]string, 0)
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		//一定是运算符
		if token == "+" || token == "-" || token == "*" || token == "/" {
			topNoe := stringToInt(stack[len(stack)-1])
			topTow := stringToInt(stack[len(stack)-2])
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, strconv.Itoa(topTow+topNoe))
			case "-":
				stack = append(stack, strconv.Itoa(topTow-topNoe))
			case "*":
				stack = append(stack, strconv.Itoa(topTow*topNoe))
			case "/":
				stack = append(stack, strconv.Itoa(topTow/topNoe))
			}
		} else { //如果是数字则直接入栈
			stack = append(stack, tokens[i])
		}
	}
	return stringToInt(stack[0])
}

func stringToInt(s string) int {
	tag, _ := strconv.Atoi(s)
	return tag
}
