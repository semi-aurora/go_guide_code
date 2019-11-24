package main

import (
	"fmt"
)

func main() {
	//将非ascii字符转为?号
	charList := "l love you ,我爱你!"
	var charEmpty string
	for _, c := range charList {
		fmt.Printf("%v\n",c)
		if c > 255 {
			charEmpty += "?"
		} else {
			charEmpty += string(c)
		}
	}
	fmt.Printf(charEmpty)
}
