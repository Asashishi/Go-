package main

import "fmt"

// 定义 Teacher 结构体
type Teacher struct {
	// 结构体对象的属性
	Name   string
	Age    int
	School string
}

func main() {

	// 实例化一个结构体对象并赋值
	var teacher0 Teacher = Teacher{"余胜军", 15, "Mykt"}
	// 获取一个结构体对象的属性
	fmt.Println(teacher0)
	fmt.Println(teacher0.Name, teacher0.Age, teacher0.School)

	// 另一种创建方法 返回指针
	var teacher1 *Teacher = &Teacher{"X", 1024, "X"}
	fmt.Println(*teacher1)
	fmt.Println(teacher1.Name, teacher1.School, teacher1.School)

	// 使用指针类型创建结构体对象实例
	var teacherPtr0 *Teacher = new(Teacher)
	// 指针类型创建的结构体赋值
	(*teacherPtr0).Name = "Java之父"
	(*teacherPtr0).Age = 23
	(*teacherPtr0).School = "Mykt"
	fmt.Println(*teacherPtr0)

	// 使用指针类型创建结构体对象实例
	var teacherPtr1 *Teacher = new(Teacher)
	// 使用此方法赋值会被编译器直接转换为指针赋值法
	teacherPtr1.Name = "CS50"
	teacherPtr1.Age = 2024
	teacherPtr1.School = "HF"
	fmt.Println(*teacherPtr1)

}
