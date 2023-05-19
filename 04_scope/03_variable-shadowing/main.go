package main

import (
	"fmt"
)

func max(x int) int {
	return x + 30
}

func main() {
	max := max(20)
	fmt.Println(max) // max is now the result, not the function
}

// don't do this; bad coding practice to shadow variables!!!
