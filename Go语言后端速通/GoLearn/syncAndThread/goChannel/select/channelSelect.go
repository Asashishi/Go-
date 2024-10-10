package main

import "fmt"

// select 进行管道的选择操作

func main() {
	// 定义一个int管道
	intChan := make(chan int, 1)
	go func() {
		intChan <- 1
	}()
	// 定义一个string管道
	strChan := make(chan string, 1)
	go func() {
		strChan <- "Asashishi"
	}()

	// 使用select case 进行管道操作带有随机性
	select {
	case v := <-intChan:
		fmt.Println(v)
	case v := <-strChan:
		fmt.Println(v)
	// 防止select阻塞
	default:
		fmt.Println("no value")
	}
}
