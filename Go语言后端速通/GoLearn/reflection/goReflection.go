package main

import (
	"fmt"
	"reflect"
)

// 反射可以在程序运行时获取变量的各种信息

// 定义一个参数为空接口的函数 使用空接口可以接收任何类型的特性
func testReflact(i interface{}) {
	// 调用反射的TypeOf函数,返回reflect.Type类型 ValueOf则返回reflect.Value
	reType := reflect.TypeOf(i)
	reValue := reflect.ValueOf(i)

	// 返回数据类型-实际类型为Type
	fmt.Println("reflect type:", reType)
	// 返回值-实际类型为value
	fmt.Println("reflect value:", reValue)

	// 从反射类型中获取值
	num := reValue.Int()
	fmt.Println("num:", num)

	// 通过接口和类型断言转回变量
	reBack := reValue.Interface()
	// 类型断言
	fmt.Println("reBack:", reBack)
	n := reBack.(int)
	fmt.Println("reBack:", n)

}

func main() {
	// 对基本数据类型进行反射
	// 定义类型
	var num int = 100
	testReflact(num)
}
