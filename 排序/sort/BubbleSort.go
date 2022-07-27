package sort

// BubbleSort 冒泡排序
type BubbleSort struct {
	BaseSort
}

func (s *BubbleSort) Sort() {
	s.bubbleSort0()
	//s.bubbleSort1()
	//s.bubbleSort2()
	//s.bubbleSort3()
}

//从前面往后面遍历
func (s *BubbleSort) bubbleSort0() {
	for i := 0; i < len(s.Array)-1; i++ { //array.length-1重复计算
		for j := 0; j < len(s.Array)-1-i; j++ { //array.length-1重复计算
			if s.cmp(j, j+1) > 0 { //array[j] > array[j + 1]
				s.swap(j, j+1)
			}
		}
	}
}

//优化一，从后面往前面遍历
func (s *BubbleSort) bubbleSort1() {
	for end := len(s.Array) - 1; end > 0; end-- { //end从最后一个元素length - 1开始遍历到1
		for begin := 1; begin <= end; begin++ { //begin从第二个元素arr[1]开始遍历到end
			if s.cmp(begin, begin-1) < 0 { //比较arr[i]和arr[i-1]的值，若arr[i-1]大，就交换
				s.swap(begin, begin-1)
			}
		}
		//一轮结束之后arr[begin]也即arr[end]即为最大值（begin == end）
	}
}

//优化二，如果序列已经完全有序，可以提前终止冒泡排序 增加sorted(缺点 对于降序的没有作用)
func (s *BubbleSort) bubbleSort2() {
	for end := len(s.Array) - 1; end > 0; end-- { //end从最后一个元素length - 1开始遍历到1
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

//优化三， 如果序列尾部已经局部有序，可以记录最后1次交换的位置，减少比较次数 增加sortedIndex
func (s *BubbleSort) bubbleSort3() {
	for end := len(s.Array) - 1; end > 0; end-- { //end从最后一个元素length - 1开始遍历到1
		sortedIndex := 1
		for begin := 1; begin <= end; begin++ { //begin从第二个元素arr[1]开始遍历到end
			if s.cmp(begin, begin-1) < 0 { //比较arr[i]和arr[i-1]的值，若arr[i-1]大，就交换
				s.swap(begin, begin-1)
				sortedIndex = begin
			}
		}
		//一轮结束之后让end等于sortedIndex
		end = sortedIndex
	}
}
