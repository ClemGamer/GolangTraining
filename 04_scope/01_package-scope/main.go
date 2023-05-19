package main

import "fmt"

// here (out of curly braces) is package level scope
// package level scope is visible anywhere of the same package
// var x is in package level scope
var x int = 10

func main() {
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println(x)
}
