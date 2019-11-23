package main

import "fmt"

func main(){
	var i1 = 5
	fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)
	//An integer: 5, it's location in memory: 0xc00000a0b8
	//地址可以存储在一个叫做指针的特殊数据类型中，在本例中这是一个指向 int 的指针
}