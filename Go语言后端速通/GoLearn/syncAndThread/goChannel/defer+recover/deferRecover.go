package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// 输出数字
func printNum() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

// 除法操作
func devide() {
	// 利用defer + recover 捕获错误
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	defer wg.Done()
	numOne := 10
	numTwo := 0
	result := numOne / numTwo
	fmt.Println(result)
}

func main() {
	wg.Add(2)
	go printNum()
	go devide()
	wg.Wait()
}
