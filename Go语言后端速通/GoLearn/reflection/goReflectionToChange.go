package main

import (
	"fmt"
	"reflect"
)

// 通过反射修可以改类型数值和结构体

// 修改基本数据类型
var num int = 100

func reflectToChangeNormal(i interface{}) {

	reValue := reflect.ValueOf(i)
	fmt.Println(reValue)

	// 使用Elem获取地址的值后使用set修改值
	reValue.Elem().SetInt(40)
}

// Person 修改结构体
type Person struct {
	Name string
	Age  int
}

// 绑定方法
func (p Person) Print() {
	fmt.Printf(p.Name, p.Age)
}
func (p *Person) SetAge(newAge int) {
	p.Age = newAge
}
func (p Person) GetAge() int {
	return p.Age
}
func (p *Person) SetName(name string) {
	p.Name = name
}
func (p Person) GetName() string {
	return p.Name
}

// 反射操作
func testPerson(i interface{}) {

	// 结构体操作方法
	val := reflect.ValueOf(i)
	fmt.Println(val)

	// 操作结构体字段

	// 获取当前字段数量
	num = val.NumField()
	fmt.Println(num)

	// 遍历取结构体内的具体值
	for i := 0; i < num; i++ {
		field := val.Field(i)
		fmt.Println(field)
	}

	num = val.NumMethod()
	fmt.Println(num)

	// 通过反射调用结构体绑定的方法 反射访问需要方法首字母大写 如果按索引访问则是按方法的ASCII码顺序 此处用方法名访问
	//Call 返回value值的切片 可以惊醒修改
	val.MethodByName("Print").Call(nil)

	// 修改字段值 修改值必须传指针地址
	//val.Elem().Field(0).SetString("AsashishiNyan")
}

func main() {

	fmt.Println(num)
	// 修改值需要将内存地址传递过去
	reflectToChangeNormal(&num)
	fmt.Println(num)

	person := Person{"Asashishi", 21}
	person.Print()

	testPerson(&person)

}
