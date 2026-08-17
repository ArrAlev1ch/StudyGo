package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan string)
	var wg sync.WaitGroup
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- "Ping"
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		msg := <-ch
		fmt.Println("Получено:", msg)
	}()
	wg.Wait()
}