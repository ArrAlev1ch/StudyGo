package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("Hello from goruutine #1")
	go fmt.Println("Hello from goruutine #2")
	go fmt.Println("Hello from goruutine #3")
	time.Sleep(time.Second)
}