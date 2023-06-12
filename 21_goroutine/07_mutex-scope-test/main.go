package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func main() {
	wg.Add(3)
	go incrementor("Foo")
	go incrementor("Bar")
	go test()
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

const times int = 100

func incrementor(s string) {
	for i := 0; i < times; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		mutex.Lock()
		x := counter
		x++
		counter = x
		fmt.Println(s, i, "Counter:", counter)
		mutex.Unlock()
	}
	wg.Done()
}

func test() {
	for i := 0; i < times; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
		counter = x
		fmt.Println("test", i, "Counter:", counter)
	}
	wg.Done()
}

// mutex 是鎖執行區間，被鎖住的區間不能有兩個func在裡面
// 所以如果有另一個func也會存取counter，還是會發生race-condition
