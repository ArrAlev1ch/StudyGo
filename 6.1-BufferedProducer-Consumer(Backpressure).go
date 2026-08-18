package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	ch := make(chan int, 5)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 20; i++ {
			if len(ch) == cap(ch) {
				fmt.Println("BUFFER FULL → producer ждет")
			}

			ch <- i
			fmt.Println("produced:", i)
		}
		
		close(ch)
	}()

	wg.Add(1)
	go func() {	
		defer wg.Done()
		for i := range ch {
			fmt.Println("consumed:", i)
			time.Sleep(time.Millisecond * 200)
		}
	}()

	wg.Wait()
}