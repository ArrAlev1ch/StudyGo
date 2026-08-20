package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	const numJobs = 20
	const numWorkers = 5

	jobs := make(chan int, 20)
	results := make(chan int, 20)

	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	wg.Wait()

	close(results)
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		job, ok := <-jobs
		if !ok {
			return
		}

		fmt.Printf("worker %d processing %d\n", id, job)

		sleepMs := rand.Intn(301) + 100
		time.Sleep(time.Duration(sleepMs) * time.Millisecond)

		results <- job

		fmt.Printf("worker %d finished %d\n", id, job)

	}	
}