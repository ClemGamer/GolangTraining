package main

import "fmt"

// ASCII(American Standard Code for Information Interchange)用8bits編碼, 可顯示字元僅限英文及部分符號.
// UTF-8(8-bit Unicode Transformation Format)是針對Unicode的可變長度編碼, 用1~4個bytes對Unicode中的所有有效編碼點進行編碼.
// UTF-8中包含ASCII, 可以說ASCII是UTF-8的子集.
func main() {
	// ASCII控制字元
	for i := 0; i < 32; i++ {
		fmt.Printf("%d\t%#[1]b\t%#[1]x\t%[1]q\n", i)
	}
	fmt.Printf("%d\t%#[1]b\t%#[1]x\t%[1]q\n", 127)

	// ASCII顯示字元
	for i := 32; i < 127; i++ {
		fmt.Printf("%d\t%#[1]b\t%#[1]x\t%[1]q\n", i)
	}
}
