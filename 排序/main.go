package main

import (
	"fmt"
	"leetcode/排序/sort"
	"leetcode/排序/tools"
)

func main() {
	array := tools.Random(10000, 1, 100)
	testSorts(array, //new InsertionSort(),
		//new InsertionSort2(),
		//new InsertionSort3(),
		//new SelectionSort(),
		//new HeapSort(),
		//new MergeSort(),
		new(sort.BubbleSort1),
		new(sort.BubbleSort2),
		//new BubbleSort3(),
		//new QuickSort(),
		//new ShellSort(),
		//new CountingSort(),
		//new RadixSort(),
		//new BucketSort()
	)
}
func testSorts(array []int, sorts ...sort.Sort) {
	for _, s := range sorts {
		newArray := tools.Copy(array)
		sort.DoSort(newArray, s)
		if !tools.IsAscOrder(newArray) {
			panic("测试未通过")
		}
	}
	sort.SortSlice(sorts)
	for _, sort := range sorts {
		fmt.Println(sort)
	}
}
