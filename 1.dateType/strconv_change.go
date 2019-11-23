package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "666"
	var an int
	var newS string
	var floatStr = "1222.3897432233333332333333333333333333333333335"
	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	//数字字符串转int
	an, _ = strconv.Atoi(orig)
	fmt.Printf("The integer is: %d\n", an)
	an = an + 5
	//int字符串转string
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
	//string转float
	nf ,_ := strconv.ParseFloat(floatStr, 64)//1222.3897432233334
	fmt.Printf("The new nf is: %v\n", nf)
	//float转string
	nfStr := strconv.FormatFloat(nf, 'f', 10, 64)
	nf2Str := strconv.FormatFloat(nf, 'e', 10, 64)
	fmt.Printf("The new nfStr is: %v\n", nfStr)
	fmt.Printf("The new nfStr is: %v\n", nf2Str)
}