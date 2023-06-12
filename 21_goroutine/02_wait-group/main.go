package main

import (
	"fmt"
	"sync"
)

// 用來等待並列執行的func完成
var wg sync.WaitGroup

func main() {

	// 說明有兩個並列執行的func要等待
	wg.Add(2)
	go foo()
	go bar()
	// 等待並列執行的func完成
	wg.Wait()

}

func foo() {
	for i := 0; i < 1000; i++ {
		fmt.Println("Foo: ", i)
	}
	// 告知WaitGroup已經完成一個func
	wg.Done()
}

func bar() {
	for i := 0; i < 1000; i++ {
		fmt.Println("Bar: ", i)
	}
	wg.Done()
}
