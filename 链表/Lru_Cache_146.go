package linkedlist

// https://leetcode.cn/problems/lru-cache
// https://leetcode.cn/problems/lru-cache-lcci/
// LRU 缓存

type LRUCache struct {
	cache      map[int]*LinkedNode
	capacity   int
	size       int
	head, tail *LinkedNode
}
type LinkedNode struct {
	key, value int
	prev, next *LinkedNode
}

func ConstructorCache(capacity int) LRUCache {
	m := make(map[int]*LinkedNode, capacity)
	dammyHead := new(LinkedNode)
	dammyTail := new(LinkedNode)
	dammyTail.prev = dammyHead
	dammyHead.next = dammyTail
	return LRUCache{m, capacity, 0, dammyHead, dammyTail}
}
func (r *LRUCache) Get(key int) int {
	cur, ok := r.cache[key]
	if !ok {
		return -1
	}
	//	取出值之后 把当前节点变成头节点
	r.moveToHead(cur)
	return cur.value
}

func (r *LRUCache) Put(key int, value int) {
	cur, ok := r.cache[key]
	if !ok { //不存在
		node := &LinkedNode{key: key, value: value}
		if r.size+1 > r.capacity { //满了 删除tail
			node := r.removeTail()
			delete(r.cache, node.key)
			r.size--
		}
		r.cache[key] = node
		//添加到头部
		r.addHead(node)
		r.size++
	} else { //存在
		cur.value = value
		//设置为新的头
		r.moveToHead(cur)
	}
}

func (r *LRUCache) moveToHead(node *LinkedNode) {
	r.removeNode(node)
	r.addHead(node)
}

func (r *LRUCache) addHead(node *LinkedNode) {
	node.next = r.head.next
	node.prev = r.head
	r.head.next.prev = node
	r.head.next = node
}

func (r *LRUCache) removeNode(node *LinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (r *LRUCache) removeTail() *LinkedNode {
	node := r.tail.prev
	r.removeNode(node)
	return node
}
