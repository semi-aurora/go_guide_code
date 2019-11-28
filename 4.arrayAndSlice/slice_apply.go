package main

import "fmt"

func main(){
	//1.从字符串生成字节切片
	s := "123sky我爱你の"
	cB := []byte(s)
	cI := []int32(s)
	cR := []rune(s)
	fmt.Printf("cB is %v\n,count is %d\ncI is %v\n,count is %d\ncR is %v\n,count is %d\n",cB,len(cB),cI,len(cI),cR,len(cR),)
	//18 还是把3个字节的汉字和日文给分开了
	//utf8.RuneCountInString(str2) utf编码
/*
   	//输出
	cB is [49 50 51 115 107 121 230 136 145 231 136 177 228 189 160 227 129 174]
	,count is 18
	cI is [49 50 51 115 107 121 25105 29233 20320 12398]
	,count is 10
	cR is [49 50 51 115 107 121 25105 29233 20320 12398]
	,count is 10
*/
	var dst []byte
	str1 := "123sky我爱你の"
	copy(dst,str1)
	fmt.Printf("%v\n",dst) //此时为[] 没有长度
	dst = make([]byte,17)
	copy(dst,str1)//超出长度17 str1长度为18 -- 不会添加长度,长度未满则为对应类型的默认值
	fmt.Printf("%v\n",dst) //此时为[49 50 51 49 50 51 49 50 51 115 107 121 230 136 145 231 136 177 228 189 160 227 129 174]

	//2.获取字符串的某一部分
	//substr := str1[1:8]//末尾在多字节字符上时,会出现乱码
	substr := str1[1:]
	fmt.Printf("%v\n",substr)//str1[1:] 时%v正常显示字符

	//3.字符串和切片的内存结构
	//双字结构,指针+字符

	//4.修改字符串中的某个字符
	s4 := "hello"
	c := []byte(s4)
	c[0] = 'c'
	s4r := string(c) // s2 == "cello"
	fmt.Printf("%v\n",s4r)

	//5.字节数组对比函数

	//6.搜索及排序切片和数组

	//7.append 函数常见操作

	//8.切片和垃圾回收
}
