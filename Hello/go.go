package main

import (
	"GoAwesomeProject/Person"
	"fmt"
)

func main() {
	fmt.Println("Hello World~")
	newPerson := Person.NewPerson("老王")
	newPerson.SetAge(10)
	newPerson.SetSal(2000)
	fmt.Println("年龄:", newPerson.GetAge(), "薪水:", newPerson.GetSal())
}
