package main

import (
	"fmt"
	"strings"
)

var str string = "Hi, I'm Marc, Hi.问问1.//。‘】【=-"
var r rune = '问'
var manyG = "gggggggggg"

func main() {
	//查找索引
	fmt.Printf("%d\n", strings.IndexRune(str, r))
	fmt.Printf("%t\n", strings.HasPrefix(str, "Hi"))
	//字符串替换
	strReplace := strings.Replace(str, "Hi", "你好", -1)
	fmt.Printf(strReplace + "\n")
	//统计字符串出现次数
	strCount := strings.Count(str, "。")
	fmt.Printf("%d\n", strCount)
	fmt.Printf("Number of double g's in %s is: ", manyG)
	fmt.Printf("%d\n", strings.Count(manyG, "gg")) //5次
	//重复字符串
	fmt.Printf(strings.Repeat(str, 2))
	//大小写
	var orig string = "Hey, how are you George?"
	var lower string
	var upper string
	fmt.Printf("The original string is: %s\n", orig)
	lower = strings.ToLower(orig)
	fmt.Printf("The lowercase string is: %s\n", lower)
	upper = strings.ToUpper(orig)
	fmt.Printf("The uppercase string is: %s\n", upper)
	//字符串修剪
	var needCutStr = "__c我爱你c  "
	var cutStr string
	var cut2Str string
	cutStr = strings.TrimSpace(needCutStr)
	cut2Str = strings.Trim(needCutStr,"_ ")//可以是多个字符串
	//TrimLeft TrimRight
	fmt.Printf("The cuted string is: %s\n",cutStr)
	fmt.Printf("The cut2Str string is: %s\n",cut2Str)
}
