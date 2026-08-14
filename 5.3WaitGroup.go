package main

import (
	"fmt"
	"sync"
	"time"
)

func WaitGroup() {
	var wg sync.WaitGroup
	wg.Add(5)

	for range 5 {
		go func() {
			defer wg.Done()
			fmt.Println("Работаю...")
		}()
		time.Sleep(time.Second)
	}
	wg.Wait()
}

func main() {
	WaitGroup()
}