package sort

// BubbleSort1 冒泡排序
type BubbleSort1 struct {
	BaseSort
}

func (s *BubbleSort1) Sort(array []int) {
	s.bubbleSort(array)
}

func (s *BubbleSort1) bubbleSort0(array []int) {
	for i := 0; i < len(array)-1; i++ { //array.length-1重复计算
		for j := 0; j < len(array)-1-i; j++ { //array.length-1重复计算
			if s.cmp(j, j+1) > 0 { //array[j] > array[j + 1]
				s.swap(j, j+1)
			}
		}
	}
}

//优化一，从后面往前面遍历
func (s *BubbleSort1) bubbleSort(array []int) {
	for end := len(array) - 1; end > 0; end-- { //end从最后一个元素length - 1开始遍历到1
		for begin := 1; begin <= end; begin++ { //begin从第二个元素arr[1]开始遍历到end
			if s.cmp(begin, begin-1) < 0 { //比较arr[i]和arr[i-1]的值，若arr[i-1]大，就交换
				s.swap(begin, begin-1)
			}
		}
		//一轮结束之后arr[begin]也即arr[end]即为最大值（begin == end）
	}
}
