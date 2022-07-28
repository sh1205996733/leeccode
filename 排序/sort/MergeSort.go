package sort

// MergeSort 归并排序
type MergeSort struct {
	BaseSort
}

func (s *MergeSort) Sort() {
	s.mergeSort()
}

// private T[] leftArray;
/**
 * 执行流程
 * ① 不断地将当前序列平均分割成2个子序列 ✓ 直到不能再分割（序列中只剩1个元素）
 * ② 不断地将2个子序列合并成一个有序序列 ✓ 直到最终只剩下1个有序序列
 */
func (s *MergeSort) mergeSort() {
	//leftArray = (T[]) new Comparable[array.length >> 1]
	s.sort(0, len(s.Array))
}

// 对 [begin, end) 范围的数据进行归并排序
func (s *MergeSort) sort(begin, end int) {
	//不断地将当前序列平均分割成2个子序列 ✓ 直到不能再分割（序列中只剩1个元素）
	if end-begin < 2 { //只剩1个元素
		return
	}
	mid := (end + begin) >> 1
	s.sort(begin, mid)
	s.sort(mid, end)
	//不断地将2个子序列合并成一个有序序列 ✓ 直到最终只剩下1个有序序列
	s.merge(begin, mid, end)
}

// 2 1  4
// 将 [begin, mid) 和 [mid, end) 范围的序列合并成一个有序序列
func (s *MergeSort) merge(begin, mid, end int) {
	//li, le := 0, mid-begin
	//ri, re := mid, end
	//ai := begin
	// 备份左边数组
	//for (int i = li; i < le; i++) {
	//	leftArray[i] = array[begin+i];
	//}
	//while (li < le) {// 如果左边还没有结束,左边结束直接返回
	//	//cmp(array[ri], leftArray[li]) <= 0 改为<=就失去稳定性
	//	if (ri < re && cmp(array[ri], leftArray[li]) < 0) {// 如果右边没有结束并且array[ri]<leftArray[li],右边先结束执行else
	//		array[ai++] = array[ri++];
	//	}else {
	//		array[ai++] = leftArray[li++];
	//	}
	//}
}
