package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param k int整型
 * @param target int整型
 * @return bool布尔型
 */
func kInArray(k int, target int) bool {
	if target == k*2+1 || target == k*3+1 {
		return true
	}
	// write code here
	return false
}
