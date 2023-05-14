package stringutil

import "fmt"

// 大寫開頭的變數在package外可見(visible / exported)
var MyName string = "ClemGamer"

// 小寫開頭的變數在package外不可見(not visible / unexported)
var notVisibleVar string = "You can't see me."

func PrintSome() {
	fmt.Println(notVisibleVar)
}
