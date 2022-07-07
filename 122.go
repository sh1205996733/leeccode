package main

import "fmt"

func main() {
	arr := []int{7, 1, 5, 5, 4, 5, 6, 4}
	fmt.Println("maxProfit-----", maxProfit(arr))
}
func maxProfit1(prices []int) int {
	len := len(prices)
	if len <= 1 {
		return 0
	}
	var max, sum int
	min := prices[0]
	for i := 1; i < len; i++ {
		price := prices[i] - min
		if price >= max {
			max = price
		} else {
			sum += max
			min = prices[i]
			max = 0
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return sum + max
}

func maxProfit(prices []int) int {
	ans := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			ans += prices[i] - prices[i-1]
		}
	}
	return ans
}
