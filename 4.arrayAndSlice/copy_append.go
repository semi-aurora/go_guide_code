package main
import "fmt"

func main() {
	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)//size 10

	//copy
	n := copy(slTo, slFrom)//slTo - 目标slice copy内容存放的容器, slFrom - 源slice copy的内容
	fmt.Println(slTo)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	//append
	sl3 := []int{1, 2, 3}
	s6 := 12
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
	for i:=0;i< len(sl3);i++  {
		fmt.Printf("%x\n",&sl3[i])
	}
	//slice内存中是连续的
	fmt.Printf("%x\n",&s6)

	//append slice
	sl11 :=[]byte{'1','2','3'}
	sl22 :=[]byte{'4','5','6','a'}
	bytes := append(sl11, sl22...)
	fmt.Printf("%v\n",bytes)

	//AppendByte 是一个append实现的类似函数
}

//to control it
func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)//原始长度
	n := m + len(data)//添加后长度
	if n > cap(slice) { // if necessary, reallocate 是否过cap超了
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2) //2倍长度添加
		copy(newSlice, slice)//把原来的复制过来 这里是指针操作
		slice = newSlice
	}
	slice = slice[0:n]//增长
	copy(slice[m:n], data)//复制
	return slice//返回
}