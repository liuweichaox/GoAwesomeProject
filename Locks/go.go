package main

import (
	"fmt"
	"github.com/go-redis/redis/v9"
	"golang.org/x/net/context"
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

func redisLock() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incr()
		}()
	}
	wg.Wait()
}

var ctx = context.Background()

func incr() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})

	var lockKey = "counter_lock"
	var counterKey = "counter"

	// lock
	resp := client.SetNX(ctx, lockKey, 1, time.Second*5)
	lockSuccess, err := resp.Result()

	if err != nil || !lockSuccess {
		fmt.Println(err, "lock result: ", lockSuccess)
		return
	}

	// counter ++
	getResp := client.Get(ctx, counterKey)
	cntValue, err := getResp.Int64()
	if err == nil {
		cntValue++
		resp := client.Set(ctx, counterKey, cntValue, 0)
		_, err := resp.Result()
		if err != nil {
			// log err
			println("set value error!")
		}
	}
	println("current counter is ", cntValue)

	delResp := client.Del(ctx, lockKey)
	unlockSuccess, err := delResp.Result()
	if err == nil && unlockSuccess > 0 {
		println("unlock success!")
	} else {
		println("unlock failed", err)
	}
}
