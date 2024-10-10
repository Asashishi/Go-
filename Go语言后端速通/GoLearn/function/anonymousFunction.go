package main

// 匿名函数 没有函数名，可以赋值给变量进行重复第哦啊有

import (
	"fmt"
)

// 将匿名函数赋值给全局变量进行调用
var func_test = func(num1, num2 int) int {
	return num1 + num2
}

func main() {
	// 创建匿名函数，并将函数的值赋给变量，此时变量为函数类型
	result := func(num1, num2 int) int {
		return num1 + num2
	} // (10, 20)直接调用

	//变量调用
	fmt.Println(result(10, 20))

	// 全局变量调用匿名函数、
	fmt.Println(func_test(10, 20))
}
