package main

import (
	"fmt"
)

// 管道 本质是一个队列
// 先进先出 线程安全 协程操作管道时不会发生资源争抢
// 管道有类型
// 管道为引用类型 必须初始化才能使用 即make之后才能使用

func main() {

	// 声明一个管道 并初始化
	var channel chan any
	channel = make(chan any, 5)

	fmt.Println(channel)
	// 将数据存入管道 但不能超过定义的容量
	channel <- 10
	str := "Asashishi"
	channel <- str

	// 取当前管道用量和实际大小
	fmt.Println(len(channel), cap(channel))

	// 从管道依次取数据
	getOne := <-channel
	getTwo := <-channel
	fmt.Println(getOne, getTwo)

	fmt.Println(len(channel), cap(channel))

	/* 没有使用协程的情况下,管道中的数据如果已经取完,再取则会报错 */
}
