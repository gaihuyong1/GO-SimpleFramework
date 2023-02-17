package main

// 请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
// 实现 LRUCache 类：
// LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
// 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行

//使用哈希表与双向链表实现

type LRUCache struct{
	size int
	capacity int
	cache map[int]*LinkNode
	head *LinkNode
	tail *LinkNode
}

type LinkNode struct{
	key,value int
	pre,next *LinkNode
}

func initLinkNode(key,value int)*LinkNode{
	return &LinkNode{
		key:key,
		value: value,
	}
}

func Constructor(capacity int)LRUCache{
	cache:=LRUCache{
		cache: map[int]*LinkNode{},
		head:initLinkNode(0,0),
		tail: initLinkNode(0,0),
		capacity: capacity,
	}
	cache.head.next=cache.tail
	cache.tail.pre=cache.head
	return cache
}

func (cache *LRUCache)Get(key int)int{
	if _,ok:=cache.cache[key];ok{
		return -1
	}
	node:=cache.cache[key]
	cache.moveToHead(node)
	return node.value
}

func (cache *LRUCache)Put(key int,value int){
	if _,ok:=cache.cache[key];!ok{
		node:=initLinkNode(key,value)
		cache.cache[key]=node
		cache.addToHead(node)
		cache.size++
		if cache.size>cache.capacity{
			removed:=cache.removeTail()
			delete(cache.cache,removed.key)
			cache.size--
		}
	}else{
		node:=cache.cache[key]
		node.value=value
		cache.moveToHead(node)
	}
}

func (cache *LRUCache)addToHead(node *LinkNode){
	node.pre=cache.head
	node.next=cache.head.next
	cache.head.next.pre=node
	cache.head.next=node
}

func (cache *LRUCache)removeNode(node *LinkNode){
	node.pre.next=node.next
	node.next.pre=node.pre
}

func (cache *LRUCache)moveToHead(node *LinkNode){
	cache.removeNode(node)
	cache.addToHead(node)
}

func (cache *LRUCache)removeTail()*LinkNode{
	node:=cache.tail.pre
	cache.removeNode(node)
	return node
}