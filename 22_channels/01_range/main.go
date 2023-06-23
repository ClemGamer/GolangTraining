package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) // 若沒有關閉通道,會導致deadlock,因為main還在等著取得下一個message,但已經沒有下一個了
	}()

	for i := range c {
		fmt.Println(i)
	}
}
