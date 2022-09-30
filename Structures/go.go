package main

import "fmt"

func main() {
	// 创建一个新的结构体
	person := Person{"liu", "wei chao 1", 1}
	fmt.Printf("%v\n", person)
	// 也可以使用 key => value 格式
	person2 := Person{first_name: "liu", last_name: "wei chao 2", id: 2}
	fmt.Printf("%+v\n", person2)
	person3 := Person{first_name: "liu"}
	// 略的字段为 0 或 空
	fmt.Printf("%#v\n", person3)

	var person4 Person
	person4.first_name = "lee"
	person4.last_name = "wei"

	fmt.Println("last_name 更改前：" + person4.last_name)
	change(&person4)
	fmt.Println("last_name 更改后：" + person4.last_name)

	fmt.Println(person4.getFullName())
}

func change(person *Person) {
	person.last_name = "han"
}

type Person struct {
	first_name string
	last_name  string
	id         int
}

func (person Person) getFullName() string {
	return person.first_name + " " + person.last_name
}
