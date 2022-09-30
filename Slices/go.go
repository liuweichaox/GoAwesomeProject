package main

import "fmt"

func main() {
	var numbers = []string{"init"}
	numbers = append(numbers, "a")
	numbers = append(numbers, "b")
	numbers = append(numbers, "c")
	for index, value := range numbers {
		fmt.Printf("%d=%v\n", index, value)
	}
}
