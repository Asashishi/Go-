package main

import "fmt"

// SayHello 接口定义 接口的方法一旦有一种绑定了结构体,就需要全部实现接口中的方法 接口不能直接创建实例!
type SayHello interface {
	// SayHello 声明方法
	SayHello()
}

/*
结构体, 自定义类型能实现多个接口,接口间也可以继承接口
接口是一个指针类型,没有被初始化就使用则会输出nil
空接口 所有类型都实现了空接口 可以吧任意变量赋值给空接口
多态 不同类型,结构体实现接口的方法可以不同
*/

// Integer 不一定要结构体去实现接口方法,自定义类型也可以
type Integer int

func (i Integer) SayHello() {
	fmt.Println("I am an Integer")
}

type Japanese struct {
}
type Chinese struct {
}

// SayHello 结构体实现接口方法,并无接口继承关键字
func (JaPerson Japanese) SayHello() {
	fmt.Println("こにちは!")
}
func (CnPerson Chinese) SayHello() {
	fmt.Println("你好!")
}

// 接收绑定SayHello方法的结构体参数直接调用受绑定的方法
func greet(Person SayHello) {
	Person.SayHello()
}

func main() {

	// 结构体调用实现方法
	CN := Chinese{}
	JA := Japanese{}
	greet(CN)
	greet(JA)

	// 自定义类型调用实现方法
	i := Integer(0)
	i.SayHello()
}
