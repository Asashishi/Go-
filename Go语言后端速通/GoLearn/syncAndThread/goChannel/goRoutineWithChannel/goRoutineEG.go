package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func routineWrite(intChan chan int) {
	defer wg.Done()
	for i := 0; i < 75; i++ {
		intChan <- i
		if i == 74 {
			close(intChan)
		}
	}

}

func routineRead(intChan chan int) {
	defer wg.Done()
	for i := range intChan {
		fmt.Println(i)
	}
}

func main() {
	// 协程共同操作管道
	// 创建一个管道
	intChan := make(chan int, 75)
	wg.Add(2)
	go routineWrite(intChan)
	go routineRead(intChan)
	wg.Wait()

}

// 当管道只写不读,或读取超限,就会出现阻塞
