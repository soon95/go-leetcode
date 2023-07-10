package main

import (
	"fmt"
)

/*
*146. LRU 缓存
https://leetcode.cn/problems/lru-cache/

通过一个哈希表和一个双向链表来解决问题
*/
type LRUCache struct {
	Dic      map[int]*CacheNode
	Capacity int

	Head       *CacheNode
	Tail       *CacheNode
	NodeLength int
}

type CacheNode struct {
	Key      int
	Val      int
	Previous *CacheNode
	Next     *CacheNode
}

func Constructor(capacity int) LRUCache {

	head := &CacheNode{
		Key: -1,
		Val: -1,
	}
	tail := &CacheNode{
		Key: -1,
		Val: -1,
	}
	head.Next = tail
	tail.Previous = head

	return LRUCache{
		Dic:        map[int]*CacheNode{},
		Capacity:   capacity,
		Head:       head,
		Tail:       tail,
		NodeLength: 0,
	}

}

func (this *LRUCache) Get(key int) int {

	node, exist := this.Dic[key]

	if !exist {
		return -1
	}

	this.Remove(node)
	this.Insert(node, this.Head)

	return node.Val
}

func (this *LRUCache) Put(key int, value int) {

	node, exist := this.Dic[key]

	if exist {
		this.Remove(node)
	}

	node = &CacheNode{
		Key: key,
		Val: value,
	}

	if this.Capacity == this.NodeLength {

		this.Remove(this.Tail.Previous)

		this.Insert(node, this.Head)

	} else {

		this.Insert(node, this.Head)
	}

}

func (this *LRUCache) Insert(node, target *CacheNode) {

	next := target.Next

	target.Next = node
	node.Previous = target

	if next != nil {
		node.Next = next
		next.Previous = node
	}

	this.Dic[node.Key] = node

	this.NodeLength++

}

func (this *LRUCache) Remove(node *CacheNode) {

	previous := node.Previous
	next := node.Next

	previous.Next = next

	if next != nil {
		next.Previous = previous
	}

	delete(this.Dic, node.Key)

	this.NodeLength--

}

func main() {

	lRUCache := Constructor(2)
	lRUCache.Put(1, 1)                  // 缓存是 {1=1}
	lRUCache.Put(2, 2)                  // 缓存是 {1=1, 2=2}
	fmt.Printf("%v\n", lRUCache.Get(1)) // 返回 1
	lRUCache.Put(3, 3)                  // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	fmt.Printf("%v\n", lRUCache.Get(2)) // 返回 -1 (未找到)
	lRUCache.Put(4, 4)                  // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	fmt.Printf("%v\n", lRUCache.Get(1)) // 返回 -1 (未找到)
	fmt.Printf("%v\n", lRUCache.Get(3)) // 返回 3
	fmt.Printf("%v\n", lRUCache.Get(4)) // 返回 4

}
