package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

// func init 會在 func main 之前先執行
func init() {
	// gomaxprocs 設定 go runtime 可以使用的 CPU 數量
	// 可使用的 CPU 數量大於 1 則 go keyword 執行的 func 會以 parallel 的方式執行
	// 但是當子執行緒多餘 CPU 數量時，多出的部分還是會以 concurrency 的方式執行
	// parallel: 不同的 func 同時在不同 CPU 上執行
	// 可使用的 CPU 數量等於 1 則 go keyword 執行的 func 會以 concurrency 的方式執行
	// concurrency: 不同的 func 輪流取得 CPU time 執行
	fmt.Println("number of cpu:", runtime.NumCPU())
	fmt.Println("previous gomaxprocs:", runtime.GOMAXPROCS(runtime.NumCPU()))
}

func main() {

	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 20; i++ {
		fmt.Println("Foo: ", i)
		time.Sleep(time.Duration(5 * time.Millisecond))
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 20; i++ {
		fmt.Println("Bar: ", i)
		time.Sleep(time.Duration(10 * time.Millisecond))
	}
	wg.Done()
}
