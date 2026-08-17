package main

import (
	"fmt"
	"sync"
)

type Store struct {
	mu sync.RWMutex
	m map[string]string
}


func (s *Store) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.m[key]
}

func (s *Store) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}

func main() {
	store := &Store{
		m: make(map[string]string),
	}
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		store.Set("key", "value")
		fmt.Println("Писатель записал: key = value")
	}()
	
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		go func() {
			defer wg.Done()
			store.Get("key")
			fmt.Printf("Читатель %d прочитал: key = value\n", i)
		}()

	}

	wg.Wait()
}