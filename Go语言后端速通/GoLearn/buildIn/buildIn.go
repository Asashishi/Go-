package main

import "fmt"

// 常用内置函数
func main() {

	str := "golang"
	// len 取字符串行数
	fmt.Println(len(str))

	// new 分配数据类型 第一个实参为类型 其返回这为类型新分配的零值指针
	num := new(int) // 为 num 分配内存 此时 num 为一个 int类型的指针
	fmt.Printf("%v %v %v \n", num, &num, *num)
}
