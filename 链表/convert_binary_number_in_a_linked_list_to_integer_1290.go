package linkedlist

//https://leetcode.cn/problems/convert-binary-number-in-a-linked-list-to-integer/
// 二进制链表转整数

// 方法一：模拟 十进制转二进制的逆过程
// 时间复杂度：O(N)
// 空间复杂度：O(N)
//5÷2=2余1
//2÷2=1余0
//1÷2=0余1
//===> 得出二进制 101 .
//反推回去 商 x 除数 + 余数
//=> 0 x 2 + 1 = 1
//-> 1 x 2 + 0 = 2
//-> 2 x 2 +1 = 5
func getDecimalValue0(head *ListNode) int {
	ans := 0
	for head != nil {
		ans <<= 1
		ans += head.Val
		//ans = ans*2 + head.Val
		head = head.Next
	}
	return ans
}

// 方法二：递归
// 时间复杂度：O(N)
// 空间复杂度：O(1)
var base = 1

func getDecimalValue(head *ListNode) int {
	if head == nil {
		return 0
	}
	ans := getDecimalValue(head.Next)
	ans = head.Val*base + ans
	base <<= 1
	return ans
}
