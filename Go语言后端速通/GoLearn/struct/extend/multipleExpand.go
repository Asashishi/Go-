package main

import "fmt"

type NoOne struct {
	One int
}

func (one NoOne) PrintOne() {
	fmt.Println(one.One)
}

type NoTwo struct {
	Two int
}

func (tow NoTwo) PrintTwo() {
	fmt.Println(tow.Two)
}

// NoThree Golang支持多继承,可以访问被继承的所有属性和方法
type NoThree struct {
	// *NoOne 使用指针也能进行继承,取值时需通过指针取变量 * 如 fmt.Println(*NoThree.NoOne)
	NoOne
	NoTwo
	Three string
}

// 结构体的字段可以是结构体类型
type NoFour struct {
	A NoThree
}

func main() {
	// 创建新的结构体实例时可以给字段指定赋值,也可直接按顺序赋值
	Test := NoThree{NoOne{1}, NoTwo{2}, "Test"}
	Test.PrintOne()
	Test.PrintTwo()
	fmt.Println(Test.Three)

	A := NoThree{NoOne{1}, NoTwo{2}, "A"}
	TestA := NoFour{A}
	fmt.Println(TestA)
}
