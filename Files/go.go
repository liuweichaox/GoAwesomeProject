package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("temp.txt")
	if err == nil {
		file.WriteString("hello world")
	}
	fmt.Println("create temp.txt")
	defer file.Close()
}
