package main

import "fmt"

func main() {
	var age int
	fmt.Print("Input your age: ")
	// ScanIn 将传入的值对应到相应变量的内存地址 类型需要与定义的类型一致(自动判定类型不过将报错)
	fmt.Scanln(&age)

	fmt.Println("Your age is:", age)
}
