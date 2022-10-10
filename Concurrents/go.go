package main

import (
	"fmt"
	"time"
)

func say(s string, c chan string) {
	fmt.Println("say start ", time.Now())
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("say end ", time.Now())
	c <- s
	close(c)
}

func main() {
	fmt.Println("start ", time.Now())
	c := make(chan string)
	go say("hello", c)
	result := <-c
	fmt.Println("end ", result, " ", time.Now())
}
