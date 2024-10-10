package main

import "fmt"

// Golang语言的继承意义在于方法复用性

// Animal 定义结构体
type Animal struct {
	Name   string
	Weight float32
}

// func (an Animal) Shout()：
// 这是一个值接收者方法，意味着方法使用接收者的副本进行操作，不能修改原始接收者的值
// 需要修改接收者的值时使用 func (an *Animal) Shout()

// Shout Speak 给Animal结构体绑定方法
func (an Animal) Shout() {
	fmt.Println("I can shout!")
}

func (an Animal) ShowInfo() {
	fmt.Println(an.Name, an.Weight)
}

// Cat Cat继承Animal
type Cat struct {
	Animal
	Type string
}

func (anCat Cat) Scratch() {
	fmt.Println("Scratch!")
}
func (anCat Cat) ShowInfo() {
	fmt.Println(anCat.Name, anCat.Weight, anCat.Type)
}

func main() {

	cat := Cat{}

	//cat := Cat{}：cat 是一个 Cat 类型的实例，可以直接使用和修改。
	//cat := &Cat{}：cat 是一个指向 Cat 实例的指针，通过指针可以访问和修改原始实例的值。

	// 设置属性
	cat.Name = "anCat"
	cat.Weight = 8.15
	cat.Type = "Cat"

	// 调用父结构体方法 两种方式
	cat.Animal.Shout()
	cat.Shout()

	/*
		注意
		当有继承关系的结构体对字段进行访问时(方法/属性),如果不指定则采用就近原则,类似与多态
		如 Animal 和 Cat 同时设置了Age属性 访问Cat.Age而不是Cat.Animal.Age时,将取Cat.Age
	*/

	// 调用房钱结构体方法
	cat.Scratch()
	cat.ShowInfo()
}

/*
结构体可以跨包访问首字母小写的方法
*/
