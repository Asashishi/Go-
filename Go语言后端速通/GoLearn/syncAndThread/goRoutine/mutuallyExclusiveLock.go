package main

import (
	"fmt"
	"sync"
)

var goNum int
var wg sync.WaitGroup

// 互斥锁 Mutex
var lock sync.Mutex

func add() {
	defer wg.Done()
	// 加解锁
	for i := 1; i <= 10000000; i++ {
		lock.Lock()
		goNum++
		lock.Unlock()
	}
}

func sub() {
	defer wg.Done()
	// 加解锁
	for i := 1; i <= 10000000; i++ {
		lock.Lock()
		goNum--
		lock.Unlock()
	}
}

func main() {
	go add()
	go sub()
	wg.Add(2)
	wg.Wait()
	fmt.Println("goNum:", goNum)
}
