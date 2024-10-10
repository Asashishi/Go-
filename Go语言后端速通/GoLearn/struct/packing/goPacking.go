package packing

import "fmt"

/*
1: 其他包下不能访问首字母小写的结构体变量与方法
2: 使用工厂函数进行初始化让包外访问
3: 提供get set方法与逻辑限制
*/

type person struct {
	name string // 其他包可以访问
	age  int    // 其他包不能访问
}

// NewPerson 定义工厂模式函数
func NewPerson(name string) *person {
	return &person{
		name: name,
	}
}

// SetAge 定义封装字段的get set函数 绑定结构体 *person
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("Type range error!")
	}
}
func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetName(name string) {
	p.name = name
}
func (p *person) GetName() string {
	return p.name
}
