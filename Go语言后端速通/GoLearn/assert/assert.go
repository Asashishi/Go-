package main

import "fmt"

type MyStruct struct{}

func checkType(value interface{}) {
	// 类型断言 及在使用某变量前判断其是否为某个结构体或可否强转为某个结构体 以便后续操作
	if _, ok := value.(MyStruct); ok {
		fmt.Println("value is of type MyStruct")
	} else {
		fmt.Println("value is NOT of type MyStruct")
	}
}

func main() {
	var x MyStruct
	checkType(x)

	var y int
	checkType(y)
}
