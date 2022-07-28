package main

import (
	"fmt"
	"leetcode/排序/sort"
	"leetcode/排序/tools"
)

func main() {
	array := tools.Random(10, 1, 100)
	testSorts(array,
		new(sort.SelectionSort),
		//new HeapSort(),
		new(sort.MergeSort),
		new(sort.BubbleSort),
		new(sort.InsertionSort),
		//new QuickSort(),
		//new ShellSort(),
		//new CountingSort(),
		//new RadixSort(),
		//new BucketSort()
	)
}
func testSorts(array []int, sorts ...sort.Sort) {
	for i, s := range sorts {
		newArray := tools.Copy(array)
		sort.DoSort(newArray, s)
		if !tools.IsAscOrder(newArray) {
			fmt.Printf("case %d 测试未通过\n", i)
			return
		}
	}
	sort.SortSlice(sorts)
	for _, sort := range sorts {
		fmt.Println(sort)
	}
}
