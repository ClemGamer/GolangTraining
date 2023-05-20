package main

import "fmt"

func main() {
	fmt.Println()

	words := "hello"
	fmt.Println("english elphabet use 1 byte")
	fmt.Println(words, []byte(words))

	fmt.Println()
	words = "你好"
	fmt.Println("chinese use 3 byes")
	fmt.Println(words, []byte(words))

	fmt.Println()
	n := 20320
	fmt.Println(n, string(n), []byte(string(n)))

	fmt.Println()
	xb := []byte{228, 189, 160}
	fmt.Println(xb, string(xb), []rune(string(xb)))

	fmt.Println()
}
