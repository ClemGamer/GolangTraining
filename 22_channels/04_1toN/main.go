package main

import "fmt"

func main() {

	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 1000; i++ {
			c <- i
		}
		close(c)
	}()

	n := 5
	for i := 0; i < n; i++ {
		go func() {
			for i := range c {
				fmt.Println(i)
			}
			done <- true
		}()
	}

	for i := 0; i < n; i++ {
		<-done
	}

}
