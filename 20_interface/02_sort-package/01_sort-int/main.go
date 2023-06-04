package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{3, 5, 1, 8, 7, 9, 3, 0}
	sort.Ints(xi)
	fmt.Println(xi)

	sort.Sort(sort.Reverse(sort.IntSlice(xi)))
	fmt.Println(xi)
}
