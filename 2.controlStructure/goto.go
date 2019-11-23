package main

import "fmt"

func main() {

	//只会循环到i=0的j=3就退出了
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				break LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
	fmt.Println()

	//j==4,j==5都没有输出
LABEL2:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL2
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
	fmt.Println()

	//j==4没有输出,但是j==5有输出
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
	fmt.Println()

	//和continue+label一样,只有j==4,j==5都没有输出
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				break
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
	fmt.Println()
}