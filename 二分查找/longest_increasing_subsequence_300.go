package binary_search

// 方法一：二分查找：准备tail数组存放最长上升子序列，核心思想就是越小的数字越要往前放，这样后面就会有更多的数字可以加入tails数组
// 只能确定数量 但不能确认具体的最长上升子序列
//
//	将nums中的数不断加入tail，当nums中的元素比tail中的最后一个大时 可以放心push进tail，
//	否则进行二分查找，让比较小的数二分查找到合适的位置，让后面有更多的数字与这个数形成上升子序列
//
// 时间复杂度：O(nlogn)
// 空间复杂度：O(n)
func lengthOfLIS(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	size := 0
	heap := make([]int, len(nums))
	for _, num := range nums {
		begin := 0
		end := size
		for begin < end { //遍历所有牌堆的牌顶 找到num在heap中的插入位置
			mid := (begin + end) >> 1
			if heap[mid] >= num { //mid左边
				end = mid
			} else {
				begin = mid + 1
			}
		}
		heap[begin] = num
		if begin == size {
			size++
		}

	}
	return size
}
