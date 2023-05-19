package main

import "fmt"

var y int = 10

func main() {
	x := 20
	fmt.Println(x)
}

// x is not visible here

func foo() {
	// x is not visible here too
}
