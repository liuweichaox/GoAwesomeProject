package main

import (
	"fmt"
)

func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200
	var ret int

	/* 调用函数并返回最大值 */
	ret = max(a, b)

	fmt.Printf("最大值是 : %d\n", ret)

	hello("刘大大")

	x, y := swap("Google", "Runoob")
	fmt.Println(x, y)

	var getSequence = func() func() int {
		i := 0
		return func() int {
			i += 1
			return i
		}
	}
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()
	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	testCallBack(1, callBack)
	testCallBack(2, func(x int) int {
		fmt.Printf("我是回调，x：%d\n", x)
		return x
	})

	var user UserInfo
	user.lastName = "Liu"
	user.firstName = "WeiChao"
	fmt.Println("name = ", user.getUserName())

	fmt.Println(getLastName(user))
}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 声明局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

/* 哈喽 */
func hello(name string) {
	fmt.Printf("hello " + name)
}

/* 包裹 */
func swap(x string, y string) (string, string) {
	return x + " a", y + " b"
}

// 声明一个函数类型
type cb func(int) int

func testCallBack(x int, f cb) {
	f(x)
}

func callBack(x int) int {
	fmt.Printf("我是回调，x：%d\n", x)
	return x
}

/* 定义结构体 */
type UserInfo struct {
	id        int32
	lastName  string
	firstName string
	age       string
	email     string
	address   string
	password  string
}

func (user UserInfo) getUserName() string {
	return user.lastName + " " + user.firstName
}

func getLastName(user UserInfo) string {
	return user.lastName
}

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
