package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	numbers := [6]int{1, 2, 3, 5}
	for idx, value := range numbers {
		fmt.Printf("第 %d 位 value 的值 = %d\n", idx, value)
	}

	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	// 读取 key 和 value
	for key, value := range map1 {
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}

	// 读取 key
	for key := range map1 {
		fmt.Printf("key is: %d\n", key)
	}

	// 读取 value
	for _, value := range map1 {
		fmt.Printf("value is: %f\n", value)
	}

	/* 定义局部变量 */
	var a int = 10
	fmt.Printf("---break---")
	/* for 循环 */
	for a < 20 {
		fmt.Printf("a 的值为 : %d\n", a)
		a++
		if a > 15 {
			/* 使用 break 语句跳出循环 */
			break
		}
	}

	a = 10
	fmt.Printf("---continue---")
	/* for 循环 */
	for a < 20 {
		if a == 15 {
			/* 跳过此次循环 */
			a = a + 1
			continue
		}
		fmt.Printf("a 的值为 : %d\n", a)
		a++
	}

	// 不使用标记
	fmt.Println("---- 不使用标记 break ----")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break
		}
	}

	// 使用标记
	fmt.Println("---- 使用标记 break label ----")
re:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break re
		}
	}

	a = 10

	/* 循环 */
LOOP:
	for a < 20 {
		if a == 15 {
			/* 跳过迭代 */
			a = a + 1
			goto LOOP
		}
		fmt.Printf("a的值为 : %d\n", a)
		a++
	}
}
