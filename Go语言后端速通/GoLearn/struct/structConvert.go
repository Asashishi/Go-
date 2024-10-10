package main

import "fmt"

type Student struct {
	Age int
}

type Person struct {
	Age int
}

func main() {

	// 结构体的转换
	var Asashishi Student = Student{Age: 21}
	var AsaNeon Person = Person{Age: 21}

	// 强制转换，需要结构体属性完全相同!
	Asashishi = Student(AsaNeon)
	fmt.Println(Asashishi.Age)

	// 使用结构体定义结构体，Stu为新的结构体类型，然后使用强制转换
	type Stu Student
	var Assassin Student = Student{Age: 21}
	var Ash Stu = Stu{Age: 21}
	Ash = Stu(Assassin)
	fmt.Println(Ash.Age)

}
