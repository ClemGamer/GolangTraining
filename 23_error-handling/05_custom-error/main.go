package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

var ErrNorgateMath = errors.New("norgate math: square root of negative number")

func main() {
	v, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(v)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// return 0, fmt.Errorf("norgate math: square root of negative number: %v", f) // Package fmt 也提供自訂錯誤訊息的函數
		// return 0, errors.New("norgate math: square root of negative number")
		return 0, ErrNorgateMath
	}
	return math.Sqrt(f), nil
}
