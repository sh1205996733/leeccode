package linkedlist

import "container/list"

// https://leetcode.cn/problems/design-hashmap/
// 设计哈希映射

type MyHashMap struct {
	data []list.List
}
type entry struct {
	key, value int
}

func ConstructorHashMap() MyHashMap {
	return MyHashMap{
		data: make([]list.List, 769),
	}
}

func (this *MyHashMap) Put(key int, value int) {
	front := this.data[this.hash(key)]
	for cur := front.Front(); cur != nil; cur = cur.Next() {
		if cur.Value.(entry).key == key {
			cur.Value = entry{key, value}
			return
		}
	}
	this.data[this.hash(key)].PushBack(entry{key, value})
}

func (this *MyHashMap) Get(key int) int {
	front := this.data[this.hash(key)]
	for cur := front.Front(); cur != nil; cur = cur.Next() {
		if cur.Value.(entry).key == key {
			return cur.Value.(entry).value
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	h := this.hash(key)
	for cur := this.data[h].Front(); cur != nil; cur = cur.Next() {
		if cur.Value.(entry).key == key {
			this.data[h].Remove(cur)
		}
	}
}

func (this *MyHashMap) hash(key int) int {
	return key % cap(this.data)
}
