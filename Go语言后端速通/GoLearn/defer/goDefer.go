package main

import "fmt"

func d_add(num1, num2 int) int {
	// defer 会在其外部函数执行完后输出 执行顺序为栈的执行顺序,后进先出
	defer fmt.Println(num1, num2)
	defer fmt.Println(num1 + num2)
	num1 += 10
	num2 += 10
	// defer压入栈中时会把变量值同时压入,在定义defer语句之后对defer影响的变量进行操作,栈中的变量不会被操作
	sum := num1 + num2
	return sum
}

func main() {
	d_add(30, 60)
	// 利用defer的延迟机制,可以在方法中进行资源回收,做时机控制
}
