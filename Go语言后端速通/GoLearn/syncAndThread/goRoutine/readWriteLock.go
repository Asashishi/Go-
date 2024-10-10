package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写锁 RWMutex 适用于读多写少 写触发锁 读不影响 读写同时发生时将产生影响

var rml sync.RWMutex
var WG sync.WaitGroup

func read() {
	defer WG.Done()
	// 启动读锁
	rml.RLock()
	fmt.Println("Start to read")
	time.Sleep(1 * time.Second)
	fmt.Println("End to read")
	// 关闭读锁
	rml.RUnlock()
}

func write() {
	defer WG.Done()
	// 启动写锁
	rml.Lock()
	fmt.Println("Start to write")
	time.Sleep(3 * time.Second)
	fmt.Println("End to write")
	// 关闭写锁
	rml.Unlock()
}

func main() {
	for i := 0; i < 15; i++ {
		WG.Add(1)
		go read()
		if i%5 == 0 {
			WG.Add(1)
			go write()
		}
	}
	WG.Wait()

}
