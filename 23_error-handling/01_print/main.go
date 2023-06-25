package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("err happened:", err)
		// err happened: open no-file.txt: no such file or directory
		log.Println("err happened:", err)
		// Package log's function Println writes message with date and time to standard error
		// 2023/06/25 14:24:00 err happened: open no-file.txt: no such file or directory
		// try "go run main.go 2&> /dev/null"
		// log.Pringln()的訊息會被丟到"/dev/null"
		// 因為0是stdin,1是stdout,2是stderr
	}
}
