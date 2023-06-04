package main

import (
	"fmt"
	"sort"
)

func main() {

	xs := []string{"Ben", "Apple", "Banana", "No", "Kaaa"}

	// asc
	sort.Strings(xs)
	fmt.Println(xs)

	// desc
	sort.Sort(sort.Reverse(sort.StringSlice(xs)))
	fmt.Println(xs)
}
