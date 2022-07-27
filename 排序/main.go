package main

import (
	"fmt"
	"leetcode/排序/sort"
	"leetcode/排序/tools"
)

func main() {
	array := tools.TailAscOrder(1, 10000, 100)
	//fmt.Println(array)
	testSorts(array, //new InsertionSort(),
		//new InsertionSort2(),
		//new InsertionSort3(),
		//new SelectionSort(),
		//new HeapSort(),
		//new MergeSort(),
		new(sort.BubbleSort),
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
