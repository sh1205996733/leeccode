package linkedlist

// https://leetcode.cn/problems/linked-list-random-node/
// 链表随机节点
import "math/rand"

type Solution struct {
	data []int
}

func ConstructorSolution(head *ListNode) Solution {
	var data []int
	for node := head; node != nil; node = node.Next {
		data = append(data, node.Val)
	}
	return Solution{data}
}

func (this *Solution) GetRandom() int {
	return this.data[rand.Intn(len(this.data))]
}
