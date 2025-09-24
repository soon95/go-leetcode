package no146

type LRUCache struct {
	size int

	nodeMap map[int]*LRUCacheNode

	nodeLinkedList *LRUCacheNodeLinkedList
}

type LRUCacheNodeLinkedList struct {
	headNode *LRUCacheNode
	tailNode *LRUCacheNode
}

// 删除某个节点
func (l *LRUCacheNodeLinkedList) delete(node *LRUCacheNode) {
	if node.pre == nil {
		l.headNode = node.next
		if l.headNode != nil {
			l.headNode.pre = nil
		}
	}
	if node.next == nil {
		l.tailNode = node.pre
		if l.tailNode != nil {
			l.tailNode.next = nil
		}
	}

	if node.pre != nil && node.next != nil {
		node.pre.next = node.next
		node.next.pre = node.pre
	}

}

// 加入头部
func (l *LRUCacheNodeLinkedList) append(node *LRUCacheNode) {

	node.pre = nil
	node.next = nil

	if l.headNode == nil && l.tailNode == nil {
		l.headNode = node
		l.tailNode = node
		return
	}
	node.next = l.headNode
	l.headNode.pre = node
	l.headNode = node
}

type LRUCacheNode struct {
	key   int
	value int
	pre   *LRUCacheNode
	next  *LRUCacheNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		size:           capacity,
		nodeMap:        make(map[int]*LRUCacheNode),
		nodeLinkedList: &LRUCacheNodeLinkedList{},
	}
}

func (this *LRUCache) Get(key int) int {
	node, exist := this.nodeMap[key]
	if exist {

		this.nodeLinkedList.delete(node)
		this.nodeLinkedList.append(node)

		return node.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {

	node, exist := this.nodeMap[key]
	if exist {
		node.value = value
		this.nodeLinkedList.delete(node)
	}
	node = &LRUCacheNode{
		key:   key,
		value: value,
	}
	if !exist && len(this.nodeMap) >= this.size {
		tailNode := this.nodeLinkedList.tailNode
		if tailNode != nil {
			delete(this.nodeMap, tailNode.key)
			this.nodeLinkedList.delete(tailNode)
		}
	}

	this.nodeMap[key] = node
	this.nodeLinkedList.append(node)
}
