package main

import (
	"fmt"
	"reflect"
)

// 使用反射获取结构体

type Student struct {
	Name string
	Age  int
}

// 使用空接口接收任何类型
func reflact(i interface{}) {

	fmt.Printf("%v\n", reflect.TypeOf(i))
	fmt.Printf("%v\n", reflect.ValueOf(i))

	// 查看强转是否成功进行断言
	j, ok := i.(Student)
	if !ok {
		fmt.Printf("%v\n", ok)
	} else {
		fmt.Printf("%v\n", j)
	}
}

func main() {
	stu := Student{"Rili", 18}
	reflact(stu)
}
