package main

import (
	"fmt"
)

type Options struct {
	par1 int
	par2 float64
	par3 string
	par4 bool
}

func main() {
	typechecking(1,1,123,"'123'",123.2,true)
	typechecking2(1,1,Options{par1:1})
}

func typechecking(a int, b int,values ... interface{}) {
	for _, value := range values {
		switch v := value.(type) {
			case int: fmt.Println(v)
			case float64: fmt.Println(v)
			case string: fmt.Println(v)
			case bool: fmt.Println(v)
			default: fmt.Println(v)
		}
	}
}

func typechecking2(a int, b int,option Options)  {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(option)
}