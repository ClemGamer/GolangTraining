package main

import (
	"fmt"
	"strconv"
)

func main() {

	n := 2
	c := make(chan string)
	done := make(chan bool) // 自己實作等待goroutine完成

	putInt := func(name string) {
		// wg.Add(1) // 若執行兩個以上可能會發生 race condition
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("%s %d", name, i)
		}
		done <- true // 在這個例子給true或false都可以,只是要有message傳進channel
	}

	for i := 0; i < n; i++ {
		go putInt(strconv.Itoa(i))
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}

}
