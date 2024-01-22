package dp

// 买卖股票的最佳时机 II
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/description/

// 方法一:动态规划
// 7, 1, 5, 3, 6, 4
// 7, 18,1, 5, 3, 6, 4
func maxProfit20(prices []int) int {
	var sum int
	var max int
	var minP = prices[0]
	for _, price := range prices {
		if max <= price-minP {
			max = price - minP
		} else {
			sum += max
			minP = price
			max = 0
		}
		if price < minP {
			minP = price
		}
	}
	return sum
}

func maxProfit2(prices []int) int {
	ans := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			ans += prices[i] - prices[i-1]
		}
	}
	return ans
}
