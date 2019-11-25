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
}
