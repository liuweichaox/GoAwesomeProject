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

	var f bool = true
	var g bool = false
	if f && g {
		fmt.Printf("第一行 - 条件为 true\n")
	}
	if f || g {
		fmt.Printf("第二行 - 条件为 true\n")
	}
	/* 修改 f 和 g 的值 */
	f = false
	g = true
	if f && g {
		fmt.Printf("第三行 - 条件为 true\n")
	} else {
		fmt.Printf("第三行 - 条件为 false\n")
	}
	if !(f && g) {
		fmt.Printf("第四行 - 条件为 true\n")
	}

	var h uint = 60 /* 60 = 0011 1100 */
	var i uint = 13 /* 13 = 0000 1101 */
	var j uint = 0

	j = h & i /* 12 = 0000 1100 */
	fmt.Printf("第一行 - j 的值为 %d\n", j)

	j = h | i /* 61 = 0011 1101 */
	fmt.Printf("第二行 - j 的值为 %d\n", j)

	j = h ^ i /* 49 = 0011 0001 */
	fmt.Printf("第三行 - j 的值为 %d\n", j)

	j = h << i /* 240 = 1111 0000 */
	fmt.Printf("第四行 - j 的值为 %d\n", j)

	j = h >> i /* 15 = 0000 1111 */
	fmt.Printf("第五行 - j 的值为 %d\n", j)

	var k int = 21
	var l int

	l = k
	fmt.Printf("第 1 行 - =  运算符实例，l 值为 = %d\n", l)

	l += k
	fmt.Printf("第 2 行 - += 运算符实例，l 值为 = %d\n", l)

	l -= k
	fmt.Printf("第 3 行 - -= 运算符实例，l 值为 = %d\n", l)

	l *= k
	fmt.Printf("第 4 行 - *= 运算符实例，l 值为 = %d\n", l)

	l /= k
	fmt.Printf("第 5 行 - /= 运算符实例，l 值为 = %d\n", l)

	l = 200

	l <<= 2
	fmt.Printf("第 6行  - <<= 运算符实例，l 值为 = %d\n", l)

	l >>= 2
	fmt.Printf("第 7 行 - >>= 运算符实例，l 值为 = %d\n", l)

	l &= 2
	fmt.Printf("第 8 行 - &= 运算符实例，l 值为 = %d\n", l)

	l ^= 2
	fmt.Printf("第 9 行 - ^= 运算符实例，l 值为 = %d\n", l)

	l |= 2
	fmt.Printf("第 10 行 - |= 运算符实例，l 值为 = %d\n", l)

	var m int = 4
	var n int32
	var o float32
	var ptr *int

	/* 运算符实例 */
	fmt.Printf("第 1 行 - m 变量类型为 = %T\n", m)
	fmt.Printf("第 2 行 - n 变量类型为 = %T\n", n)
	fmt.Printf("第 3 行 - o 变量类型为 = %T\n", o)

	/*  & 和 * 运算符实例 */
	ptr = &m /* 'ptr' 包含了 'm' 变量的地址 */
	fmt.Printf("m 的值为  %d\n", m)
	fmt.Printf("*ptr 为 %d\n", *ptr)
}
