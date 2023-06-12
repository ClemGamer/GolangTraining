package main

import "fmt"

func main() {
	// go keyword 執行子執行緒
	// concurrency 或 parallel 要看 runtime 的 CPU 數量
	go foo()
	go bar()
	// 不會印出任何訊息
	// 因為主執行緒(main)在丟出子執行緒執行後會繼續執行
	// 執行到底就直接結束主執行緒(main)，子執行緒也一同結束
}

func foo() {
	for i := 0; i < 50; i++ {
		fmt.Println("Foo: ", i)
	}
}

func bar() {
	for i := 0; i < 50; i++ {
		fmt.Println("Bar: ", i)
	}
}
