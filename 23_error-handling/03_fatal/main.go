package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Fatalln("err happened:", err)
		// fatal calls os.Exit(1) after writing the log message
	}
}
