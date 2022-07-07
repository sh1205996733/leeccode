package main

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	cur := len(nums1) - 1
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[cur] = nums1[i]
			i--
		} else { //nums1[i] <= nums2[j]
			nums1[cur] = nums2[j]
			j--
		}
		cur--
	}
}
