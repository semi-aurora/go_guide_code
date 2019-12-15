package main

import "fmt"

type pp struct {
	name string
	age int
	sex bool
}
func(self *pp) grow(year int){
	self.age += year
}
func(self pp) nameChange(newName string){
	self.name = newName
}

func main(){
	var p pp
	p.grow(10)
	fmt.Println(p)
	p.nameChange("mi-autumn")//不会改变name
	fmt.Println(p)
	/*
		{ 10 false}
		{ 10 false}
	*/
}