package sort

import "fmt"

// SelectionSort 选择排序
type SelectionSort struct {
	BaseSort
}

func (s *SelectionSort) Sort() {
	//s.selectionSort1()
	//s.selectionSort2()
	//s.selectionSort3()
	s.selectionSort4()
}

func (s *SelectionSort) selectionSort1() { //1 2 3 4 5 6
	for end := 0; end < len(s.Array)-1; end++ {
		min := end
		for begin := end + 1; begin < len(s.Array); begin++ {
			if s.cmp(begin, min) < 0 {
				min = begin
			}
		}
		s.swap(end, min)
	}
	fmt.Println(s.Array)
}
func (s *SelectionSort) selectionSort2() { //6 5 4 3 2 1
	for end := 0; end < len(s.Array)-1; end++ {
		max := 0
		for begin := 1; begin < len(s.Array)-end; begin++ {
			if s.cmp(begin, max) > 0 {
				max = begin
			}
		}
		s.swap(max, len(s.Array)-1-end)
	}
}

func (s *SelectionSort) selectionSort3() { //6 5 4 3 2 1
	for end := len(s.Array) - 1; end > 0; end-- {
		for begin := 0; begin < end; begin++ {
			if s.cmp(begin, end) > 0 {
				s.swap(begin, end)
			}
		}
	}
}

func (s *SelectionSort) selectionSort4() { //6 5 4 3 2 1
	for end := len(s.Array) - 1; end > 0; end-- {
		max := 0
		for begin := 1; begin <= end; begin++ {
			if s.cmp(begin, max) > 0 {
				max = begin
			}
		}
		s.swap(max, end)
	}
}
