package main

import (
	"fmt"
	"time"
)

func main(){
	start := time.Now()
	longCalculation()
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}
func longCalculation(){
	time.Sleep(10000)
}