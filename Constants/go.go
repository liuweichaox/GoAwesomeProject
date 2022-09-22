package main

import "fmt"

func main() {
	//常量的定义格式：
	//const identifier [type] = value
	//显式类型定义
	const a string = "abc"
	fmt.Println("a", a)

	//隐式类型定义
	const b = "abc"
	fmt.Println("b", b)

	//多个相同类型的声明可以简写为：
	//const c_name1, c_name2 = value1, value2
	const c, d, e = 1, false, "str" //多重赋值
	fmt.Println("c", c, "d", d, "e", e)
	//常量还可以用作枚举：
	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)
	fmt.Println(Unknown, Female, Male)
}
