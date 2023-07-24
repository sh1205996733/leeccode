package monotonicstack

// https://leetcode.cn/problems/next-greater-element-i/
// https://leetcode.cn/problems/next-greater-element-ii/
// 下一个更大元素

// 单调栈 + 哈希表
// 时间复杂度O(m+n)
// 空间复杂度O(n)
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	//遍历nums2 将nums2放入一个单调递减栈中
	var stack []int
	m := make(map[int]int, 0)
	for _, num := range nums2 {
		m[num] = -1                                       //默认初始化 -1
		for len(stack) > 0 && stack[len(stack)-1] < num { //当前元素值大于栈顶元素值，弹出栈顶元素
			m[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}
	var ans []int
	for _, num := range nums1 {
		ans = append(ans, m[num])
	}
	return ans
}

// 暴力法
func nextGreaterElement0(nums1, nums2 []int) []int {
	m, n := len(nums1), len(nums2)
	res := make([]int, m)
	for i, num := range nums1 {
		j := 0
		for j < n && nums2[j] != num {
			j++
		}
		k := j + 1
		for k < n && nums2[k] < nums2[j] {
			k++
		}
		if k < n {
			res[i] = nums2[k]
		} else {
			res[i] = -1
		}
	}
	return res
}

// 下一个更大元素 II 遍历nums 单调栈中保存的是下标，从栈底到栈顶的下标在数组 nums 中对应的值是单调不升的。
// 时间复杂度O(n)
// 空间复杂度O(n)
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	var stack []int
	for i := 0; i < n*2-1; i++ {
		if i < n {
			ans[i] = -1 //默认初始化 -1
		}
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] { //当前元素值大于栈顶元素值，弹出栈顶元素
			ans[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n) //将索引和值都存到栈中
	}
	return ans
}

// https://leetcode.cn/problems/final-prices-with-a-special-discount-in-a-shop/
// 商品折扣后的最终价格 找第一个比他小的数
func finalPrices(prices []int) []int {
	ans := make([]int, len(prices))
	var stack []int
	for i, price := range prices {
		ans[i] = price
		for len(stack) > 0 && prices[stack[len(stack)-1]] >= price {
			ans[stack[len(stack)-1]] = prices[stack[len(stack)-1]] - price
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}
