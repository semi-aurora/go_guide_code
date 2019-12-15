package main

import "fmt"

type A struct {
	ax, ay int
}

type B struct {//内嵌结构体 -- 用于实现类似的集成 组合
	A
	bx, by float32
}

func main() {
	b := B{A{1, 2}, 3.0, 4.0}
	fmt.Println(b.ax, b.ay, b.bx, b.by)
	fmt.Println(b.A)
	/*
		1 2 3 4
		{1 2}
	 */
}