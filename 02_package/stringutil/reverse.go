package stringutil

// 大寫開頭的方法在packae外可見(visible / exported).
// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	return reverseTwo(s)
}
