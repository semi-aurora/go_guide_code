package main

import (
	"bytes"
	"fmt"
	"runtime"
	"sync"
)

type Info struct {
	name string
	address string
	age int
	lk sync.Mutex
}

//共享缓冲器
type SyncedBuffer struct {
	lock    sync.Mutex
	buffer  bytes.Buffer
}

func main(){
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	sayHello()
	sayWorld()
}

func sayHello(){
	for i := 0; i < 10; i++{
		fmt.Print("hello ")
		runtime.Gosched()
	}
}

func sayWorld(){
	for i := 0; i < 10; i++ {
		fmt.Println("world")
		runtime.Gosched()
	}
}

func Update(info *Info) {
	info.lk.Lock()
	// critical section:
	info.name = "123"// new value
	info.address = "12333"// new value
	info.age = 1123// new value
	// end critical section
	info.lk.Unlock()
}
