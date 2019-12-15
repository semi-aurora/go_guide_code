package main

import (
	"fmt"
	"strconv"
)

//定义一个加法函数，传入channel类型，每计算一次，quit加1
func Add(x, y int, quit chan int) {
	z := x + y
	fmt.Println(z)
	//写数据
	quit <- 1
}

//接受数据，并赋值给我
func Read(ch chan int) {
	value := <-ch
	fmt.Println("value" + strconv.Itoa(value))
}

func Write(ch chan int) {
	//ch <- 10
}

func main() {
	//定义一个channel类型切片数组
	chs := make([]chan int, 10)
	//循环执行加法函数
	for i := 0; i < 10; i++ {
		//拿取一个channel
		chs[i] = make(chan int)
		//执行goroutine和发送一个channel类型数据
		go Add(i, i, chs[i])
	}
	for _, v := range chs {
		//接口channel类型数据
		  fmt.Println(v)
	}
}
