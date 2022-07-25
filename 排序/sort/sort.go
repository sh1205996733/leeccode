package sort

import (
	"fmt"
	"reflect"
	"time"
)

type BaseSort struct {
	Name      string
	Array     []int
	CmpCount  int
	SwapCount int
	Time      int64
}

func DoSort(array []int, sort Sort) {
	if array == nil || len(array) < 2 {
		return
	}
	f := reflect.ValueOf(sort).Elem()
	f.FieldByName("Name").SetString(f.Type().Name())
	f.FieldByName("Array").Set(reflect.ValueOf(array))
	begin := time.Now()
	sort.Sort(array)
	f.FieldByName("Time").SetInt(time.Since(begin).Milliseconds())
}

type Sort interface {
	Sort([]int)
}

//func (a *BaseSort) Len() int { // 重写 Len() 方法
//	return len(a)
//}
//func (a *BaseSort) Swap(i, j int) { // 重写 Swap() 方法
//	a[i], a[j] = a[j], a[i]
//}
//func (a *BaseSort) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
//	return a[j].CreatedAt.UnixNano() < a[i].CreatedAt.UnixNano()
//}

/**
 * 返回值等于0，代表 array[i1] == array[i2]
 * 返回值小于0，代表 array[i1] < array[i2]
 * 返回值大于0，代表 array[i1] > array[i2]
 */
func (i *BaseSort) cmp(i1, i2 int) int {
	i.CmpCount++
	return i.Array[i1] - (i.Array[i2])
}

func (i *BaseSort) swap(i1, i2 int) {
	i.SwapCount++
	i.Array[i1], i.Array[i2] = i.Array[i2], i.Array[i1]
}

func SortSlice(array []Sort) {
	for end := len(array) - 1; end > 0; end-- { //end从最后一个元素length - 1开始遍历到1
		for begin := 1; begin <= end; begin++ { //begin从第二个元素arr[1]开始遍历到end
			if CompareTo(array[begin], array[begin-1]) < 0 { //比较arr[i]和arr[i-1]的值，若arr[i-1]大，就交换
				array[begin], array[begin-1] = array[begin-1], array[begin]
			}
		}
	}
}
func CompareTo(i, j interface{}) int64 {
	oi := reflect.ValueOf(i).Elem()
	oj := reflect.ValueOf(j).Elem()

	result := oi.FieldByName("Time").Int() - oj.FieldByName("Time").Int()
	if result != 0 {
		return result
	}

	result = oi.FieldByName("CmpCount").Int() - oj.FieldByName("CmpCount").Int()
	if result != 0 {
		return result
	}

	return oi.FieldByName("SwapCount").Int() - oj.FieldByName("SwapCount").Int()
}
func (i *BaseSort) String() string {
	timeStr := fmt.Sprintf("耗时：%ds(%dms)", i.Time/1000, i.Time)
	compareCountStr := "比较：" + numberString(i.CmpCount)
	swapCountStr := "交换：" + numberString(i.SwapCount)
	stableStr := fmt.Sprintf("稳定性：%v", i.isStable())
	return "【" + i.Name + "】\n" +
		stableStr + " \t" +
		timeStr + " \t" +
		compareCountStr + "\t " +
		swapCountStr + "\n" +
		"------------------------------------------------------------------"
}

func numberString(number int) string {
	if number < 10000 {
		return fmt.Sprintf("%d", number)
	}
	if number < 100000000 {
		return fmt.Sprintf("%d万", number/10000.0)
	}
	return fmt.Sprintf("%d亿", number/100000000.0)
}

func (i *BaseSort) isStable() bool {
	//if (this instanceof BucketSort) return true;
	//if (this instanceof RadixSort) return true;
	//if (this instanceof CountingSort) return true;
	//if (this instanceof ShellSort) return false;
	//if (this instanceof SelectionSort) return false;
	//Student[] students = new Student[20];
	//for (int i = 0; i < students.length; i++) {
	//	students[i] = new Student(i * 10, 10);
	//}
	//sort((T[]) students);
	//for (int i = 1; i < students.length; i++) {
	//	int score = students[i].score;
	//	int prevScore = students[i - 1].score;
	//	if (score != prevScore + 10) return false;
	//}
	return true
}
