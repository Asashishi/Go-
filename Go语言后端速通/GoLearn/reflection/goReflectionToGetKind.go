package main

import (
	"fmt"
	"reflect"
)

// 使用反射获取变量类别

type Japanese struct {
	Name string
	Age  int
}

func GetKind(i interface{}) {

	// .Kind()方法获取变量类别
	Kind := reflect.TypeOf(i).Kind()
	fmt.Println(Kind)
	Value := reflect.ValueOf(i).Kind()
	fmt.Println(Value)

	// 类型和类别的区别 类别: 大类别类似Integer 类型: 更为具体的类型
	// 类型断言转换成结构体
	i, ok := i.(Japanese)
	if !ok {
		fmt.Println(ok)
	} else {
		fmt.Println(i)
	}
}

func main() {
	Rui := Japanese{"Rui", 30}
	GetKind(Rui)
}
