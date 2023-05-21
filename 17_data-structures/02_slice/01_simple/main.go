package main

import "fmt"

func main() {

	xi := []int{} // xi is defined without number in [] is a slice

	fmt.Println(xi)
	fmt.Println(len(xi))

	xi = append(xi, 1) // append 1 at last of xi and return a new slice

	fmt.Println(xi)
	fmt.Println(len(xi))

}

// 雖然slice可以用append動態延展，但這麼做會重新分配記憶體，效率不佳
// 在使用slice時最好還是用make預先分配好要用的大小
// 需要延展時，再用make預先分配更大的大小預留空間，避免頻繁的重新分配記憶體
