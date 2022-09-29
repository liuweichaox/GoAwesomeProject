package main

import "fmt"

func main() {
	//第一种，指定变量类型，如果没有初始化，则变量默认为零值。

	var a string = "Runoob"
	fmt.Println("a", a)

	var b, c int = 1, 2
	fmt.Println("b", b, "c", c)

	// 字符串为 ""（空字符串）
	var d string
	fmt.Println("d", d)

	// 没有初始化就为零值
	var e int
	fmt.Println("e", e)

	// bool 零值为 false
	var f bool
	fmt.Println("f", f)

	var h *int
	fmt.Println("h", h)

	var i []int
	fmt.Println("i", i)

	var g map[string]int
	fmt.Println("g", g)

	var k chan int
	fmt.Println("k", k)

	var l func(string) int
	fmt.Println("l", l)

	var m error // error 是接口]
	fmt.Println("m", m)

	//第二种，根据值自行判定变量类型。
	var n = true
	fmt.Println("n", n)

	//第三种，如果变量已经使用 var 声明过了，再使用 := 声明变量，就产生编译错误，格式：v_name := value
	//var intVal int
	//intVal := 1 // 这时候会产生编译错误，因为 intVal 已经声明，不需要重新声明
	o := "Runoob" // var f string = "Runoob"
	fmt.Println("o", o)

	//多变量声明
	//类型相同多个变量, 非全局变量
	//var vname1, vname2, vname3 type
	//var vname1, vname2, vname3 = v1, v2, v3 // 和 python 很像,不需要显示声明类型，自动推断
	//vname1, vname2, vname3 := v1, v2, v3 // 出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误
	// 这种因式分解关键字的写法一般用于声明全局变量
	//var (
	//	vname1 v_type1
	//	vname2 v_type2
	//)
	var p, q int

	fmt.Println("p", p, "q", q)

	var r, s = 1, "abc"
	fmt.Println("r", r, "s", s)

	t, u := 3, "hello"
	fmt.Println("t", t, "u", u)

	//多变量可以在同一行进行赋值
	var v, w int
	var x string
	v, w, x = 5, 7, "abc"
	fmt.Println("v", v, "w", w, "x", x)
}
