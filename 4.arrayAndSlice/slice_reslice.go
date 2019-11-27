package main

import "fmt"

func main(){
	var arr2 = []int{1,2,3,4,5,6,7}
	slice := arr2[3:3]
	slice1 := arr2[3:4]
	fmt.Printf("len slice is : %v ,the slice1 is %v ,and value slice is %v",len(slice),len(slice1),slice)
}
