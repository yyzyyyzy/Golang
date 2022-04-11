package main

import "container/list"

// lru 中的数据
type Node struct {
	K, V interface{}
}

// 链表 + map
type LRU struct {
	list     *list.List
	cacheMap map[interface{}]*list.Element
	Size     int
}

// 初始化一个LRU
func NewLRU(cap int) *LRU {
	return &LRU{
		Size:     cap,
		list:     list.New(),
		cacheMap: make(map[interface{}]*list.Element, cap),
	}
}

// 获取LRU中数据
func (lru *LRU) Get(k interface{}) (v interface{}, ret bool) {
	// 如果存在，则把数据放到链表最前面
	if ele, ok := lru.cacheMap[k]; ok {
		lru.list.MoveToFront(ele)
		return ele.Value.(*Node).V, true
	}

	return nil, false
}

// 设置LRU中数据
func (lru *LRU) Set(k, v interface{}) {
	// 如果存在，则把数据放到最前面
	if ele, ok := lru.cacheMap[k]; ok {
		lru.list.MoveToFront(ele)
		ele.Value.(*Node).V = v // 更新数据值
		return
	}

	// 如果数据是满的，先删除数据，后插入
	if lru.list.Len() == lru.Size {
		last := lru.list.Back()
		node := last.Value.(*Node)
		delete(lru.cacheMap, node.K)
		lru.list.Remove(last)
	}

	ele := lru.list.PushFront(&Node{K: k, V: v})
	lru.cacheMap[k] = ele
}
