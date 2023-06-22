package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var count int64
var wg sync.WaitGroup

func increment(name string) {
	for i := 1; i <= 20; i++ {
		time.Sleep(time.Duration(5 * time.Millisecond))
		atomic.AddInt64(&count, 1)
		// fmt.Println(name, count)
		// 用atomic可以避免race condition
		// 但接著若要印出count,因為不是在atomic內保證
		// 所以還是會檢查到race condition且看到不同的過程
		// 但count++的運算結果還是被atomic保證正確的
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go increment("foo")
	go increment("bar")
	wg.Wait()
	fmt.Println(count)
}
