package main

import "fmt"

func main(){
	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}
	var arr1 = new([5]int)
	var arr2[5]int
	fmt.Printf("%T\n",arr1)//*[5]int
	fmt.Printf("%T\n",arr2)//[5]int

	//prove array copy where array being assigned
	var arr3[3]int
	fmt.Printf("%x\n",&arr3)
	fmt.Printf("%x\n",&arr3[0])
	fmt.Printf("%x\n",&arr3[1])
	fmt.Printf("%x\n",&arr3[2])

	//arr3 = [3]int{1,2,3} //same with pointer
	arr3[0] = 1
	arr3[1] = 2
	arr3[2] = 3
	//same with pointer

	fmt.Printf("%x\n",&arr3)
	fmt.Printf("%x\n",&arr3[0])
	fmt.Printf("%x\n",&arr3[1])
	fmt.Printf("%x\n",&arr3[2])
	//same 数组赋值时内存地址不变 只有值发生改变


	var arrAge = [5]int{18, 20, 15, 22}//可以少于5个
	var arrLazy = [...]int{5, 6, 7, 8, 22}
	// var arrLazy = []int{5, 6, 7, 8, 22}//不写...也是ok的 严格来说是变成了切片的语法 可实际上还是[]int
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	//var arrKeyValue = []string{3: "Chris", 4: "Ron"}
	for i := 0; i < len(arrAge); i++ {
		fmt.Printf("Age at %d is %d\n", i, arrAge[i])
	}
	fmt.Println()

	for i := 0; i < len(arrLazy); i++ {
		fmt.Printf("Number at %d is %d\n", i, arrLazy[i])
	}
	fmt.Printf("\n")

	for i := 0; i < len(arrKeyValue); i++ {
		fmt.Printf("Person at %d is %s\n", i, arrKeyValue[i])
	}

	var arrLazy2 = []int{5, 6, 7, 8, 22}
	fmt.Printf("%T",arrLazy2)
}
