package main

import "fmt"

func main() {

	// 声明并初始化管道
	var intChan chan int
	intChan = make(chan int, 5)

	intChan <- 10
	// 关闭管道 关闭后不能写入数据,但是可以读取
	close(intChan)

	num := <-intChan
	fmt.Println(num)

}
