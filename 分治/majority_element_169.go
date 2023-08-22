package divide_conquer

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

/*
	https://leetcode.cn/problems/majority-element/solution/tong-su-yi-dong-mo-er-tou-piao-fa-by-ni-h4m1b/

“同归于尽消杀法” ：

由于多数超过50%, 比如100个数，那么多数至少51个，剩下少数是49个。

第一个到来的士兵，直接插上自己阵营的旗帜占领这块高地，此时领主 winner 就是这个阵营的人，现存兵力 count = 1。

如果新来的士兵和前一个士兵是同一阵营，则集合起来占领高地，领主不变，winner 依然是当前这个士兵所属阵营，现存兵力 count++；

如果新来到的士兵不是同一阵营，则前方阵营派一个士兵和它同归于尽。 此时前方阵营兵力count --。（即使双方都死光，这块高地的旗帜 winner 依然不变，因为已经没有活着的士兵可以去换上自己的新旗帜）

当下一个士兵到来，发现前方阵营已经没有兵力，新士兵就成了领主，winner 变成这个士兵所属阵营的旗帜，现存兵力 count ++。

就这样各路军阀一直以这种以一敌一同归于尽的方式厮杀下去，直到少数阵营都死光，那么最后剩下的几个必然属于多数阵营，winner 就是多数阵营。（多数阵营 51个，少数阵营只有49个，死剩下的2个就是多数阵营的人）
*/
func majorityElement1(nums []int) int {
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

// 方法三：分治 不断从数组的中间进行递归分割，直到每个区间的个数是1，然后向上合并左右区间个数较多的数，向上返回。
// 复杂度分析：时间复杂度：O(nlogn)，不断二分，复杂度是logn，二分之后每个区间需要线性统计left与right的个数，复杂度是n。空间复杂度：O(logn)，递归栈的消耗，不断二分。
// 时间复杂度：O(NlogN)
// 空间复杂度：O(logN)
func majorityElement(nums []int) int {
	return getMode(nums, 0, len(nums)-1)
}

func getMode(nums []int, le int, ri int) int {
	// 不断从数组中间分割 直到每个区间的个数是1
	if le == ri {
		return nums[le]
	}
	// 取中间值
	mid := (ri-le)/2 + le

	// 向上合并左右区间个数较多的数
	left := getMode(nums, le, mid)
	right := getMode(nums, mid+1, ri)
	if left == right {
		return left
	}
	leftCount := getCount(nums, left, le, ri)
	rightCount := getCount(nums, right, le, ri)
	if leftCount > rightCount {
		return left
	}
	return right
}

func getCount(nums []int, target int, le int, ri int) int {
	var count int
	for i := le; i < ri; i++ {
		if nums[i] == target {
			count++
		}
	}
	return count
}
