package binary_search

// https://leetcode.cn/problems/median-of-two-sorted-arrays/

// 方法一：暴力法：合并两个有序数组。
// 时间复杂度：O(n+m)
// 空间复杂度：O(n+m)
func findMedianSortedArrays0(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	var newArray []int
	// 合并数组
	first, second := 0, 0
	for first < len1 && second < len2 {
		if nums1[first] <= nums2[second] {
			newArray = append(newArray, nums1[first])
			first++
		} else {
			newArray = append(newArray, nums2[second])
			second++
		}
	}
	if first < len1 {
		newArray = append(newArray, nums1[first:]...)
	}
	if second < len2 {
		newArray = append(newArray, nums2[second:]...)
	}
	if len(newArray) == 0 {
		return 0
	}
	mid := len(newArray) >> 1
	if len(newArray)&1 == 1 {
		return float64(newArray[mid])
	}
	return float64(newArray[mid-1]+newArray[mid]) / 2.0
}

// 优化
func findMedianSortedArrays00(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	if len1 == 0 && len2 == 0 {
		return 0
	}
	newArray := make([]int, len1+len2)
	li, ri := 0, 0
	for cur := 0; cur < len(newArray); cur++ {
		if li < len1 && nums1[li] < nums2[ri] {
			newArray[cur] = nums1[li]
			li++
		} else { // li == len1 || nums1[li] >= nums2[ri]
			newArray[cur] = nums2[ri]
			ri++
		}
	}
	mid := len(newArray) >> 1
	if len(newArray)&1 == 1 {
		return float64(newArray[mid])
	}
	return float64(newArray[mid-1]+newArray[mid]) / 2.0
}

// 看不懂？？？ todo
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	midIndex := totalLength >> 1
	if midIndex&1 == 1 {
		return float64(getKthElement(nums1, nums2, midIndex+1))
	}
	midIndex1, midIndex2 := midIndex-1, midIndex
	return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
	return 0
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
