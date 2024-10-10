package main

import (
	"fmt"
	"sync"
)

// Goroutine 异步等待 WaitGroup *WaitGroup Add/Done/Wait

func main() {
	// 定义异步等待计数器
	var wg sync.WaitGroup
	// 启动十个协程
	for i := 1; i <= 10; i++ {
		// 等待组加一 ,.Add方法需要与需要执行的协程数一致
		wg.Add(1)
		// 协程启动匿名函数
		go func(n int) {
			// 执行完成后进行等待组减一
			defer wg.Done()
			fmt.Printf("%v times\n", n)
		}(i)
	}
	// 阻塞主线程等待异步完成
	wg.Wait()
}
