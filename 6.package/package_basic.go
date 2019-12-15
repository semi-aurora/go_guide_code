package main

import (
	"container/list"
)

func main(){
	//遍历链表
	var l list.List
	l = list.List{}
	for e := l.Front(); e != nil; e = e.Next() {
		//do something with e.Value
	}
}
