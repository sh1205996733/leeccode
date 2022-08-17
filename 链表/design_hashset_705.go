package linkedlist

import "container/list"

// https://leetcode.cn/problems/design-hashset/
// 设计哈希集合

// 方法一 数组
// 时间复杂度为 O(1)
// 空间复杂度：O(n)
//type MyHashSet struct {
//	data []int
//}
//
//func Constructor() MyHashSet {
//
//	return MyHashSet{
//		data: []int{0: -1, 1000000: 0},
//	}
//}
//
//func (this *MyHashSet) Add(key int) {
//	this.data[key] = key
//}
//
//func (this *MyHashSet) Remove(key int) {
//	this.data[key] = -1
//}
//
//func (this *MyHashSet) Contains(key int) bool {
//	return this.data[key] == key
//}

// MyHashSet 方法二 使用链表
// 时间复杂度为 O(n)
// 空间复杂度：O(1)
type MyHashSet struct {
	data []list.List
}

func Constructor() MyHashSet {
	return MyHashSet{
		data: make([]list.List, 769),
	}
}

func (this *MyHashSet) Add(key int) {
	if this.Contains(key) {
		return
	}
	// 存在使用链地址法
	hash := this.hash(key)
	this.data[hash].PushBack(key)
}

func (this *MyHashSet) hash(key int) int {
	return key % cap(this.data)
}

func (this *MyHashSet) Remove(key int) {
	h := this.hash(key)
	for cur := this.data[h].Front(); cur != nil; cur = cur.Next() {
		if cur.Value.(int) == key {
			this.data[h].Remove(cur)
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	head := this.data[this.hash(key)]
	for cur := head.Front(); cur != nil; cur = cur.Next() {
		if cur.Value.(int) == key {
			return true
		}
	}
	return false
}
