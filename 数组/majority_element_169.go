package array

// 多数元素
// https://leetcode.cn/problems/majority-element/

// 方法一：哈希表 (遍历数组，将出现次数放到map中 再对map进行遍历)
// 时间复杂度：O(N)
// 空间复杂度：O(n)
func majorityElement0(nums []int) int {
	m := make(map[int]int, 0)
	for _, num := range nums {
		m[num]++
	}
	target := len(nums) / 2
	for k, v := range m {
		// fmt.Println(k, v)
		// 多数元素
		if v > target {
			return k
		}
	}
	return 0
}

// 方法二：摩尔投票法 Boyer-Moore 投票算法 遍历数组 初始化count = 1 遇到相同的count++ 遇到不同的count--
// 时间复杂度：O(N)
// 空间复杂度：O(1)

/* https://leetcode.cn/problems/majority-element/solution/tong-su-yi-dong-mo-er-tou-piao-fa-by-ni-h4m1b/
“同归于尽消杀法” ：

由于多数超过50%, 比如100个数，那么多数至少51个，剩下少数是49个。

第一个到来的士兵，直接插上自己阵营的旗帜占领这块高地，此时领主 winner 就是这个阵营的人，现存兵力 count = 1。

如果新来的士兵和前一个士兵是同一阵营，则集合起来占领高地，领主不变，winner 依然是当前这个士兵所属阵营，现存兵力 count++；

如果新来到的士兵不是同一阵营，则前方阵营派一个士兵和它同归于尽。 此时前方阵营兵力count --。（即使双方都死光，这块高地的旗帜 winner 依然不变，因为已经没有活着的士兵可以去换上自己的新旗帜）

当下一个士兵到来，发现前方阵营已经没有兵力，新士兵就成了领主，winner 变成这个士兵所属阵营的旗帜，现存兵力 count ++。

就这样各路军阀一直以这种以一敌一同归于尽的方式厮杀下去，直到少数阵营都死光，那么最后剩下的几个必然属于多数阵营，winner 就是多数阵营。（多数阵营 51个，少数阵营只有49个，死剩下的2个就是多数阵营的人）
*/
func majorityElement(nums []int) int {
	cnt, max := 0, 0
	for i := 0; i < len(nums); i++ {
		if cnt == 0 {
			max = nums[i]
		}
		if max == nums[i] {
			cnt++
		} else {
			cnt--
		}
	}
	return max
}
