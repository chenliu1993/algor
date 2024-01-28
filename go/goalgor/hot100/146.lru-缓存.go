/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 */

// @lc code=start
type LRUCache struct {
	capacity   int
	store      map[int]*LinkedList
	head, tail *LinkedList
	size       int
}

type LinkedList struct {
	prev, next *LinkedList
	key, value int
}

func NewLinkedList(key, value int) *LinkedList {
	return &LinkedList{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		store:    map[int]*LinkedList{},
		head:     NewLinkedList(0, 0),
		tail:     NewLinkedList(0, 0),
		size:     0,
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (this *LRUCache) moveToHead(node *LinkedList) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeNode(node *LinkedList) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
func (this *LRUCache) addToHead(node *LinkedList) {
	node.next = this.head.next
	node.prev = this.head
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.store[key]; ok {
		node := this.store[key]
		this.moveToHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) removeTail() *LinkedList {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.store[key]; ok {
		node := this.store[key]
		node.value = value
		this.moveToHead(node)
		return
	}

	node := NewLinkedList(key, value)
	this.store[key] = node
	this.addToHead(node)
	this.size++
	if this.size > this.capacity {
		removed := this.removeTail()
		delete(this.store, removed.key)
		this.size--
	}

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

