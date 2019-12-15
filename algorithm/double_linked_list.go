package main

//定义节点
type DNode struct {
	data interface{}
	prev *DNode
	next *DNode
}
//定义 双向链表
type DList struct {
	size uint64
	head *DNode
	tail *DNode
}

//初始化链表Init
func (dList *DList) Init() {
	_dlist := *(dList)
	_dlist.size = 0
	_dlist.head = nil
	_dlist.tail = nil
}

