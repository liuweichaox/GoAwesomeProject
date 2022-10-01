package main

import "fmt"

func main() {
	var i int = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, foo(uint64(i)))
}
func foo(i uint64) (result uint64) {
	if i > 0 {
		var next uint64 = i - 1
		result = i * foo(next)
		return result
	}
	return 1
}
