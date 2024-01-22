package dp

// 买卖股票的最佳时机
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/

// 方法一:动态规划
// 7, 1, 5, 3, 6, 4
// 7, 18,1, 5, 3, 6, 4
// 7,6,4,3,1
func maxProfit(prices []int) int {
	var max int
	var minP = prices[0]
	for _, price := range prices {
		if max < price-minP {
			max = price - minP
		}
		if price < minP {
			minP = price
		}
	}
	return max
}

// 方法一:单调栈
func maxProfit0(prices []int) int {
	var stack []int
	var max int
	for _, price := range prices {
		for len(stack) > 0 && price < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 || price > stack[len(stack)-1] {
			stack = append(stack, price)
		}
		if len(stack) > 0 && max < stack[len(stack)-1]-stack[0] {
			max = stack[len(stack)-1] - stack[0]
		}
	}
	return max
}
