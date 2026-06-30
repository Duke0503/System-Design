package main

import (
	"fmt"
	"sync"
	"time"
)

type Store struct {
	mu   sync.Mutex
	data map[string]int
}

func (s *Store) Set(key string, val int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = val
}

func (s *Store) Get(key string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.data[key]
}

func main() {
	store := &Store{data: make(map[string]int)}

	fmt.Println("=== Day 2 Probe - State and Where It Lives ===")
	fmt.Println("Two users. One in-memory store. What happens?")

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			store.Set("user_a_visits", i)
			fmt.Printf("[User A] recorded visit #%d\n", i)
			time.Sleep(10 * time.Millisecond)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			store.Set("user_b_visits", i)
			fmt.Printf("[User B] recorded visit #%d\n", i)
			time.Sleep(15 * time.Millisecond)
		}
	}()
	wg.Wait()

	fmt.Println("\n=== Final State (in memory) ===")
	fmt.Printf("user_a_visits: %d\n", store.Get("user_a_visits"))
	fmt.Printf("user_b_visits: %d\n", store.Get("user_b_visits"))

	fmt.Println("\n[!] Process is about to exit.")
	fmt.Println("[!] All state above will be gone.")
	fmt.Println("[?] Where should this data actually live?")
}
