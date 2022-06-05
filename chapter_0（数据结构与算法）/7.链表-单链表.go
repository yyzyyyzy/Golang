package main

//1.实现单链表的结点
type SingleLinkedListNode struct { //单链表的结点
	value interface{}
	next  *SingleLinkedListNode
}

func NewSingleLinkedListNode(data interface{}) *SingleLinkedListNode { //创建一个结点
	return &SingleLinkedListNode{
		value: data,
		next:  nil,
	}
}

func (node *SingleLinkedListNode) Value() interface{} { //返回数据
	return node.Value
}

func (node *SingleLinkedListNode) Next() interface{} { //返回结点
	return node.Next
}

//2.实现单链表
type SingleLink interface {
	//单链表的增删改查
	GetFirstNode() *SingleLinkedListNode
	InsertNodeFront(node *SingleLinkedListNode)   //头部插入
	InsertNodeBack(node *SingleLinkedListNode)    //尾部插入
	GetNodeIndex(index int) *SingleLinkedListNode //根据索引抓取指定位置的结点
	DeleteNode(dest *SingleLinkedListNode) bool   //删除一个结点
	DeleteAtIndex(index int) bool                 //删除指定索引的结点
	String() string                               //返回链表的字符串
}

type SingleLinkedList struct {
	head   *SingleLinkedListNode //链表的头指针
	length int                   //链表的长度
}

func NewSingleLinkedList() *SingleLinkedList {
	head := NewSingleLinkedListNode(nil)
	return &SingleLinkedList{
		head:   head,
		length: 0,
	}
}

func (s SingleLinkedList) GetFirstNode() *SingleLinkedListNode {
	return s.head.next
}

func (s SingleLinkedList) InsertNodeFront(node *SingleLinkedListNode) {
	if s.head == nil {
		s.head = node
		node.next = nil
		s.length++
	} else {
		bak := s.head //备份
		node.next = bak.next
		bak.next = node
		s.length++
	}
}

func (s SingleLinkedList) InsertNodeBack(node *SingleLinkedListNode) {
	//TODO implement me
	panic("implement me")
}

func (s SingleLinkedList) GetNodeIndex(index int) *SingleLinkedListNode {
	//TODO implement me
	panic("implement me")
}

func (s SingleLinkedList) DeleteNode(dest *SingleLinkedListNode) bool {
	//TODO implement me
	panic("implement me")
}

func (s SingleLinkedList) DeleteAtIndex(index int) bool {
	//TODO implement me
	panic("implement me")
}

func (s SingleLinkedList) String() string {
	//TODO implement me
	panic("implement me")
}
