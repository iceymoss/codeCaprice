package greedy

import "math"

/*
输入：prices = [7,1,5,3,6,4]
输出：7
解释：在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4 。
     随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6 - 3 = 3 。
     总利润为 4 + 3 = 7 。
*/

func maxProfit(prices []int) int {
	max := math.MinInt
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			max += (prices[i+1] - prices[i])
		}
	}
	return max
}
