package main

import (
	"fmt"
	"time"
)

func main() {
	fast := make(chan string)
	slow := make(chan string)
	
	go func() {
		fast <- "fast"
	}()
	go func() {
		time.Sleep(time.Second)
		slow <- "slow"
	}()

	select {
	case msg := <-fast:
		fmt.Println(msg)

	case msg := <-slow:
		fmt.Println(msg)
	}
}