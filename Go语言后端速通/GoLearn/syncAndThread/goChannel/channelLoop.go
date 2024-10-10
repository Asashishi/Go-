package main

import "fmt"

func main() {

	var loopChan chan int
	loopChan = make(chan int, 100)

	// 向管道中放入值
	for i := 0; i < 100; i++ {
		loopChan <- i
	}

	close(loopChan)

	// 遍历管道必须使用for-range循环 且需要在遍历前关闭管道
	for j := range loopChan {
		fmt.Println(j)
	}
}
