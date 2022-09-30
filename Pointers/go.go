package main

import "fmt"

const MAX int = 3

func main() {
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	var ptr *int

	fmt.Printf("ptr 的值为 : %x\n", ptr)

	arr := []int{10, 100, 200}
	var i int

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, arr[i])
	}

	var g int
	var gptr *int
	var gpptr **int

	g = 3000

	/* 指针 ptr 地址 */
	gptr = &g

	/* 指向指针 ptr 地址 */
	gpptr = &gptr

	/* 获取 pptr 的值 */
	fmt.Printf("变量 g = %d\n", g)
	fmt.Printf("指针变量 *gpptr = %d\n", *gptr)
	fmt.Printf("指向指针的指针变量 **gpptr = %d\n", **gpptr)

	/* 定义局部变量 */
	var x int = 100
	var y int = 200

	fmt.Printf("交换前 x 的值 : %d\n", x)
	fmt.Printf("交换前 y 的值 : %d\n", y)

	/* 调用函数用于交换值
	 * &x 指向 x 变量的地址
	 * &y 指向 y 变量的地址
	 */
	swap(&x, &y)

	fmt.Printf("交换后 x 的值 : %d\n", x)
	fmt.Printf("交换后 y 的值 : %d\n", y)
}
func swap(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址的值 */
	*x = *y   /* 将 y 赋值给 x */
	*y = temp /* 将 temp 赋值给 y */
}
