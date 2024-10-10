package main

import "fmt"

func main() {
	/*
	* 指针类型变量接收到的只能为地址值
	* 指针类型的变量指向的地址必须与指向的类型匹配
	* 基本数据类型都有对应的指针类型 (* + 类型)
	 */

	var age int = 18
	// & + 变量名 根据变量取指针
	fmt.Println(&age)

	/*
	* 定义一个指针变量 ptr 指向 int 类型
	* &age 为变量 age 的地址，是 ptr 的值
	* 此时 ptr 为地址
	* ptr 是指针类型本身也有地址
	 */
	var ptr *int = &age
	fmt.Println(ptr)  // 通过指针取 age 的地址
	fmt.Println(&ptr) // 获取 ptr 本身的地址

	// 获取 ptr 指向的地址的数据 * + 变量
	fmt.Println(*ptr)

}
