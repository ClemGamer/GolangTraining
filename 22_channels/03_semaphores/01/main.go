package main

import (
	"fmt"
)

func main() {

	c := make(chan string)
	done := make(chan bool) // 自己實作等待goroutine完成

	putInt := func(name string) {
		// wg.Add(1) // 若執行兩個以上可能會發生 race condition
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("%s %d", name, i)
		}
		done <- true // 在這個例子給true或false都可以,只是要有message傳進channel
	}

	go putInt("foo")
	go putInt("bar")

	go func() {
		<-done
		<-done // 有2個goroutine要等
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}

}
