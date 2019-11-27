package main

import "fmt"

func main(){
	//for-range循环
	items := [...]int{10, 20, 30, 40, 50}
	for _, item := range items {
		item *= 2 //值不变
	}
	fmt.Printf("%v\n",items)
	for index := range items {
		items[index] *= 2 //值改变
	}
	fmt.Printf("%v\n",items)
	//多维数组切片
	var screen [10][10]int
	for row := range screen {
		for column := range screen[row] {
			screen[row][column] = 1
		}
	}
	fmt.Printf("%v\n",screen)
}