package sort

// BubbleSort2 冒泡排序
type BubbleSort2 struct {
	BaseSort
}

func (s *BubbleSort2) Sort(array []int) {
	s.bubbleSort(array)
}

//冒泡排序
//优化二，如果序列已经完全有序，可以提前终止冒泡排序 增加sorted(缺点 对于降序的没有作用)
func (s *BubbleSort2) bubbleSort(array []int) {
	for end := len(array) - 1; end > 0; end-- { //end从最后一个元素length - 1开始遍历到1
		sorted := true
		for begin := 1; begin <= end; begin++ { //begin从第二个元素arr[1]开始遍历到end
			if s.cmp(begin, begin-1) < 0 { //比较arr[i]和arr[i-1]的值，若arr[i-1]大，就交换
				s.swap(begin, begin-1)
				sorted = false
			}
		}
		//一轮结束之后arr[begin]也即arr[end]即为最大值（begin == end）
		if sorted {
			break
		}
	}
}
