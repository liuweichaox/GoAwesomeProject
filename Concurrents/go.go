package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string, c chan string) {
	fmt.Println("say start ", time.Now())
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("say end ", time.Now())
	c <- s
	close(c)
}

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)
}

func main() {
	fmt.Println("start ", time.Now())
	c := make(chan string)
	go say("hello", c)
	result := <-c
	fmt.Println("end ", result, " ", time.Now())

	//启动多个goroutine
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
	
}
