package Person

import (
	"fmt"
)

// 定义一个结构体
type person struct {
	Name string
	age  int
	sal  float64
}

// 编写工厂模式函数
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

// 编写get和set方法
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("请输入正确的年龄")
	}
}
func (p *person) GetAge() int {
	return p.age
}
func (p *person) SetSal(sal float64) {
	if sal > 0 {
		p.sal = sal
	} else {
		fmt.Println("请输入正确的薪水")
	}
}
func (p *person) GetSal() float64 {
	return p.sal
}
