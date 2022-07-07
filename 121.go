package main

import "fmt"

func main() {
	arr := []int{1, 1, 2}
	fmt.Println("maxProfit-----", maxProfit(arr))
}
func maxProfit(prices []int) int {
	len := len(prices)
	if len <= 1 {
		return 0
	}
	var max int
	min := prices[0]
	for i := 1; i < len; i++ {
		if prices[i]-min > max {
			max = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return max
}
