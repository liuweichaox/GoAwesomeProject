package main

import "fmt"

func main() {
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c)
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c)
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c)
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c)
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c)
	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a)
	a = 21 // 为了方便测试，a 这里重新赋值为 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a)

	var d int = 21
	var e int = 10

	if d == e {
		fmt.Printf("第一行 - d 等于 e\n")
	} else {
		fmt.Printf("第一行 - d 不等于 e\n")
	}
	if d < e {
		fmt.Printf("第二行 - d 小于 e\n")
	} else {
		fmt.Printf("第二行 - d 不小于 e\n")
	}

	if d > e {
		fmt.Printf("第三行 - d 大于 e\n")
	} else {
		fmt.Printf("第三行 - d 不大于 e\n")
	}

	d = 5
	e = 20
	if d <= e {
		fmt.Printf("第四行 - d 小于等于 e\n")
	}
	if e >= d {
		fmt.Printf("第五行 - e 大于等于 d\n")
	}
}
