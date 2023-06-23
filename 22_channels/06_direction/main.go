package main

import "fmt"

func main() {
	inc := incrementer(10)
	for i := range puller(inc) {
		fmt.Println(i)
	}
}

func incrementer(n int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

// "<-chan"代表只能取出
func puller(c <-chan int) <-chan int {
	// 建立可取出也可放入的channel
	out := make(chan int)
	go func() {
		sum := 0
		for i := range c {
			sum += i
		}
		out <- sum
		close(out)
	}()
	// 回傳只可取出的channel
	return out
}
