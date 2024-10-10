package main

import (
	"GoLearn/struct/packing"
	"fmt"
)

func main() {
	// 使用封装设置值,取值
	p := packing.NewPerson("Asashishi")
	//fmt.Println(p.Name) 无法访问
	p.SetAge(10)
	fmt.Println(p.GetName(), p.GetAge())
	// 使用指针也可访问,但只能整体访问不能访问特定属性
	fmt.Println(*p)
}
