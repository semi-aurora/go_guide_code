package main

import "fmt"

type integer int
func add(a integer) integer{
	return 1 + a
}

func main(){
	var i int
	i = 2
	fmt.Println(add(integer(i)))
}