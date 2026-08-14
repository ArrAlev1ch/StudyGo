package main

import (
	"fmt"
	"sync"
)

func Ping(c chan string) {
	c <- "Ping"
}

func main() { 
	ch := make(chan string)
	var wg sync.WaitGroup
	
	wg.Add(1)
	
	go Ping(ch) 
	
	go func() {
		defer wg.Done()
		fmt.Println("Получено:", <-ch)
	}()
	
	wg.Wait()
}