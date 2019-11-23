package main

import (
	"fmt"
)

func main() {
	//指针参数的使用
	n := 0
	reply := &n
	Multiply(10, 5, reply)
	fmt.Println("Multiply:", *reply) // Multiply: 50
	fmt.Println("n:", n)             // Multiply: 50
	//多返回值(命令返回值方式)
	var min, max int
	min, max = MinMax(78, 65)
	fmt.Printf("Minmium is: %d, Maximum is: %d\n", min, max)
}

// this function changes reply:
func Multiply(a, b int, reply *int) {
	*reply = a * b
}

//命令返回值方式的函数
func MinMax(a int, b int) (min int, max int) {
	if a < b {
		min = a
		max = b
	} else { // a = b or a < b
		min = b
		max = a
	}
	return
}
