package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	c := make(chan string)

	putInt := func(name string) {
		// wg.Add(1) // 若執行兩個以上可能會發生 race condition
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("%s %d", name, i)
		}
		wg.Done()
	}

	wg.Add(2)
	go putInt("foo")
	go putInt("bar")

	go func() {
		wg.Wait()
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}

}
