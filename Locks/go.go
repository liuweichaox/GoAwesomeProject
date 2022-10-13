package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"sync"
	"time"
)

// 全局变量
var counter int

func main() {
	fmt.Println("unlock----------------------")
	unlock()

	fmt.Println("lock----------------------")
	lock()

	fmt.Println("redisLock----------------------")
	redisLock()
}

func unlock() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}

	wg.Wait()
	println(counter)
}
func lock() {
	counter = 0
	var wg sync.WaitGroup
	var l sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			l.Lock()
			counter++
			l.Unlock()
		}()
	}

	wg.Wait()
	println(counter)
}

var ctx = context.Background()

func redisLock() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
	// SET key value EX 10 NX
	set, err := rdb.SetNX(ctx, "lockKey", "value", 10*time.Second).Result()
	fmt.Println(set, err)
}
