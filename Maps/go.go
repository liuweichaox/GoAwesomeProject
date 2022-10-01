package main

import "fmt"

func main() {
	var dictMap map[string]int
	dictMap = make(map[string]int)
	dictMap["a"] = 1
	dictMap["b"] = 2
	dictMap["c"] = 3
	fmt.Println(dictMap["a"])
	fmt.Println(dictMap["b"])
	fmt.Println(dictMap["c"])

	for k, v := range dictMap {
		fmt.Printf("dictMap[%s]=%x\n", k, v)
	}
}
