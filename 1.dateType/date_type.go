package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"
)
var (
	HOME = os.Getenv("HOME")
	USER = os.Getenv("USER")
	GOROOT = os.Getenv("GOROOT")
)

func main() {
	//goos := runtime.GOOS
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)

	fmt.Printf("Home is %s\n", HOME)
	fmt.Printf("User is %s\n", USER)
	fmt.Printf("GOROOT is %s\n", GOROOT)

	timers := int64(time.Now().Nanosecond())
	rand.Seed(timers)
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d / ", a)
	}
	for i := 0; i < 5; i++ {
		r := rand.Intn(100)
		fmt.Printf("%d / ", r)
	}
	fmt.Println()

	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f / ", 100*rand.Float32())
	}
}