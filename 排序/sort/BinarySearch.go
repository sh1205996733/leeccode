package sort

// 二分搜索

// IndexOf 查找v在有序数组array中的位置
func IndexOf(array []int, v int) int {
	if array == nil || len(array) == 0 {
		return -1
	}
	begin, end := 0, len(array)
	for begin < end {
		mid := (begin + end) >> 1
		if v < array[mid] {
			end = mid
		} else if v > array[mid] {
			begin = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// Search 查找v在有序数组array中待插入位置
func Search(array []int, v int) int {
	if array == nil || len(array) == 0 {
		return -1
	}
	begin, end := 0, len(array)
	for begin < end {
		mid := (begin + end) >> 1
		if v < array[mid] {
			end = mid
		} else {
			begin = mid + 1
		}
	}
	return begin
}
