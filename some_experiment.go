package main

import "fmt"

func main() {
	var floatTest float64 = -0.11111
	fmt.Printf("%b\n",floatTest)

	for i := 0; i<7 ; i++ {
		if i%2 == 0 { continue }
		fmt.Println("Odd:", i)
	}

	var v1, v2 , v3 = "string", 3, 4
	fmt.Printf("%v %v %v\n",v1,v2,v3)
}