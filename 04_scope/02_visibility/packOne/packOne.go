package packOne

import "fmt"

// var y in package main is not visible in this package

func boo() {
	// error
	// fmt.Println(y)

	// var myName in file name.go is visible because in same package
	fmt.Println(myName)
}
