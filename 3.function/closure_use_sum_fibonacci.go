package main

import (
	"fmt"
	"time"
)

//fibonacci algorithm change to closure
func main() {
	fmt.Printf("请输入fibonacci数")
	var nNumber int
	_, _ = fmt.Scanln(&nNumber)
	start := time.Now()
	fmt.Printf("%v\n",getFibonacciValue(nNumber))
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

func getFibonacciValue(n int) int {
	sum := 0
	fib := fibonacciFunc()
	for i := 0; i <= n;i++ {
		//sum += fib() //计算综合
		sum = fib()//计算数
	}
	return sum
}

func fibonacciFunc() func() int {
	a := - 1 //保存a
	b := 1 //保存b
	//序号 -  - - 1 2 3 4 5 6 7  8 ...
	//数值 -1 1 0 1 1 2 3 5 8 13 21 ...
	return func() int {
		a, b = b, a + b
		return a + b
	}
}
