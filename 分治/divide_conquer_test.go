package divide_conquer

import (
	"fmt"
	binarytree "leetcode/二叉树"
	"math"
	"testing"
)

func TestMyPow(t *testing.T) {
	x := 0.0
	n := -2
	fmt.Println(myPow(x, n))
	fmt.Println(math.Pow(x, float64(n)))
}

func TestMajorityElement(t *testing.T) {
	nums := []int{1, 2, 2, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
}

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

func TestRangeSumBST(t *testing.T) {
	nums := []int{-10, -3, 0, 5, 9}
	low, high := 5, 15
	root := binarytree.MakePerfectBinaryTree(nums)
	fmt.Printf("%+v\n", rangeSumBST(root, low, high))
}

// 位1的个数
func TestHammingWeight(t *testing.T) {
	var nums uint32 = 4294967293
	fmt.Println(hammingWeight(nums))
}

// 数组中的第K个最大元素
func TestFindKthLargest(t *testing.T) {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	fmt.Println(findKthLargest(nums, k)) // 4
}
