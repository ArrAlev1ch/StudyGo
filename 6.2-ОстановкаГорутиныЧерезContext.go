package main

import (
	"context"
	"log"
	"time"
	"sync"
)

var wg sync.WaitGroup

func worker(ctx context.Context) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			log.Println("context canceled, stopping")
			return
		default:
			log.Println("working...")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go worker(ctx)

	time.Sleep(1 * time.Second)
	cancel()

	wg.Wait()
}