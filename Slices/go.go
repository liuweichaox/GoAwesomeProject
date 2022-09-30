package main

import "fmt"

func main() {
	var numbers []string
	numbers = append(numbers, "a")
	numbers = append(numbers, "b")
	numbers = append(numbers, "c")
	for key, value := range numbers {
		fmt.Printf("%d=%v\n", key, value)
	}
}
