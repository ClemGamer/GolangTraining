package main

import "fmt"

func main() {

	xi := make([]int, 0, 10) // 預先配置10個int大小的記憶體

	for i := 0; i < 100; i++ {
		xi = append(xi, i)
		fmt.Printf("value:%v len:%v cap:%v\n", i, len(xi), cap(xi))
	}

	fmt.Println("xi", xi)
}

// make可以預先配置一部分的記憶體給slice使用，稱為slice的容量(capacity)，避免頻繁的重新配置記憶體
// 當append會超過slice的容量時，會以當前slice的容量的兩倍重新配置記憶體
// slice是referenc type，slice的capacity其實就是slice管理的一個array的長度，slice的len是可以存取這個array的長度
