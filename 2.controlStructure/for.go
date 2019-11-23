package main

import "fmt"

func main() {
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for ix :=0; ix < len(str); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str[ix])
	}
	str2 := "chinese:日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for ix :=0; ix < len(str2); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str2[ix])
	}
	str3 := "chinese:日本語"
	fmt.Println("index int(rune) rune    char bytes")
	for index, rune_ := range str3 {
		fmt.Printf("%-2d      %d      %U '%c' % x\n", index, rune_, rune_, rune_, []byte(string(rune_)))
	}
	//print 会出现意想不到的问题 个人觉得是内部自行使用了go启动了携程
	//str := "G"
	//for i := 1; i <= 25; i++ {
	//	print(str)
	//	str += "G"
	//}
	str1 := "G"
	for i := 1; i <= 25; i++ {
		fmt.Println(str1)
		str1 += "G"
	}
}