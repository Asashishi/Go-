package main

import "fmt"

func main() {

	var name string
	var age int
	fmt.Print("Please input your name and age (split by [space]): ")

	//Scanf 格式化传入的值对应到相应变量的内存地址 类型需要与定义的类型一致(自动判定类型不过将报错)
	fmt.Scanf("%v %v", &name, &age)

	fmt.Printf("Hello %v, %v\n", name, age)
}
