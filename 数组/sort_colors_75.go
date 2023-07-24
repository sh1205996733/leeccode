package array

// 颜色分类
// https://leetcode.cn/problems/sort-colors/

// 方法一：排序
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func sortColors0(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
	}
}

// 方法二：双指针(三指针) 0，1，2 排序。一次遍历，如果是0，则移动到表头，如果是2，则移动到表尾，不用考虑1
// 时间复杂度：O(n)
// 空间复杂度：O(1)

// 1 0 2 2 0 1 -->
// 2,0,2,1,1,0 --> 0,0,1,1,2,2
func sortColors1(nums []int) {
	//all in [0,p0] 0
	//all in (p0,p2) 1
	//all in [p2,len(nums)-1] 2
	var i int
	p0, p2 := -1, len(nums)
	for i < p2 {
		if nums[i] == 0 {
			p0++
			nums[p0], nums[i] = nums[i], nums[p0]
			i++
		} else if nums[i] == 1 {
			i++
		} else {
			p2--
			nums[p2], nums[i] = nums[i], nums[p2]
		}
	}
}

func sortColors(nums []int) {
	//all in [0,p0) 0
	//all in [p0,i) 1
	//all in [p2,len(nums)-1) 2
	var i int
	p0, p2 := 0, len(nums)-1
	for i <= p2 {
		if nums[i] == 0 {
			nums[p0], nums[i] = nums[i], nums[p0]
			p0++
			i++
		} else if nums[i] == 1 {
			i++
		} else {
			nums[p2], nums[i] = nums[i], nums[p2]
			p2--
		}
	}
}
