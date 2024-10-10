package main

import "fmt"

func get_sum() func(int) int {
	/*
		闭包函数
		外层函数 get_sum 无参数传递 返回类型为一个函数 int 类型传参 int 类型返回
		内置函数 将传入的参数加上内层函数外的变量后返回
		本质为使用一个有名函数包裹一个匿名函数，并在整个匿名函数中引用匿名函数外的变量
	*/
	var sum int = 0 // 此变量一直保存在内存中，是闭包函数引入的变量,包外无法使用此变量
	return func(num int) int {
		sum += num
		return sum
	}
}

func main() {
	f := get_sum() // 将匿名函数的内层函数传给 f 此时 f 的类型为函数
	fmt.Println(f(1))
	fmt.Println(f(14)) // 调用并打印匿名函数

	// 使用场景: 需要保留或记录每次引用函数时的都某个变量时
}

// eg
// 带接收者的函数 func(this *User) 传入的接收者能使用参数调用,并能够修改其中的值
// func (this *User) ListOlineUser(msg string)
