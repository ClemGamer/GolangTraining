package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	defer func() {
		fmt.Println("first")
	}()
	defer func() {
		fmt.Println("second")
	}()
	if err != nil {
		log.Panicln("err happened:", err)
		// log.Panic calls builtin panic after printing the log message
		// builtin panic stops current goroutine and prints panic message and paniced line after running defer function in reversed order
		// panic calls os.Exit(2) to exit main thread
		// panic(err)
	}

}
