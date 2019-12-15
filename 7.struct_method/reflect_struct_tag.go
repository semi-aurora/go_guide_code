package main

import (
	"fmt"
	"reflect"
)

type TagType struct { // tags
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

func main() {
	tt := TagType{true, "Barak Obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Println(ixField.Name)//field1   			field2					field3
	fmt.Println(ixField.PkgPath)//main				main					main
	fmt.Println(ixField.Type)//bool					string					int
	fmt.Println(ixField.Tag)//An important answer	The name of the thing	How much there are
	fmt.Println(ixField.Offset)//0					8						24
	fmt.Println(ixField.Index)//[0]					[1]						[2]
	fmt.Println(ixField.Anonymous)//false			false					false
}