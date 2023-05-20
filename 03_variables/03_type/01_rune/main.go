package main

import "fmt"

func main() {
	fmt.Println()

	for i := 0; i < 128; i++ {
		fmt.Printf("%q ", i)
	}
	fmt.Println()
	fmt.Println()

	// rune is alias of int32
	// because UTF-8 uses 1~4 bytes (8~32 bits)
	fmt.Println("rune is alias of int32")
	var xr []rune = []rune{'a', 'b', 'c'}
	for _, r := range xr {
		fmt.Printf("%q\t%v\n", r, r)
	}
	fmt.Println()
}
