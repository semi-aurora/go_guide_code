// big.go
package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	// Here are some calculations with bigInts:
	im := big.NewInt(math.MaxInt64)
	in := im
	io := big.NewInt(1956)
	ip := big.NewInt(1)
	ip.Mul(im, in).Add(ip, im).Div(ip, io)
	fmt.Printf("Big Int: %v\n", ip)
	// Here are some calculations with bigInts:
	rm := big.NewRat(math.MaxInt64, 1956)
	rn := big.NewRat(-1956, math.MaxInt64)
	ro := big.NewRat(19, 56)
	rp := big.NewRat(1111, 2222)
	rq := big.NewRat(1, 1)
	rq.Mul(rm, rn).Add(rq, ro).Mul(rq, rp)
	fmt.Printf("Big Rat: %v\n", rq)

	rq2 := big.NewRat(1, 1)
	//sub subtraction 减法
	//add addition 加法
	//div division method 除法
	//mul multiplication 乘法
	//z.method(x,y) 将x method y赋值给z
	rq2.Mul(rm, rn)
	fmt.Printf("mul res: %v\n",rq2)

	/* Output:
		Big Int: 43492122561469640008497075573153004
		Big Rat: -37/112
		mul res: -1/1
	*/
}

