package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"sync"

	_ "modernc.org/sqlite"
)

const (
	goroutines = 10
	increments = 100
	expected   = goroutines * increments
)

// fileIncrement does a read → modify → write on a JSON counter file.
// No locking, no coordination. This is what "just use a file" looks like.
func fileIncrement(path string) {
	data, _ := os.ReadFile(path)
	var count int
	json.Unmarshal(data, &count)
	count++
	out, _ := json.Marshal(count)
	os.WriteFile(path, out, 0644)
}

func fileCounter() int {
	path := "/tmp/day3_counter.json"
	os.WriteFile(path, []byte("0"), 0644)
	defer os.Remove(path)

	var wg sync.WaitGroup
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < increments; j++ {
				fileIncrement(path)
			}
		}()
	}
	wg.Wait()

	data, _ := os.ReadFile(path)
	var result int
	json.Unmarshal(data, &result)
	return result
}

// sqliteCounter does the same increments but through SQLite.
// SQLite serializes writes — one writer at a time, atomically.
// SetMaxOpenConns(1) ensures all goroutines share one connection
// (and one in-memory database) so SQLite's serialization is visible.
func sqliteCounter() int {
	db, _ := sql.Open("sqlite", ":memory:")
	db.SetMaxOpenConns(1)
	defer db.Close()

	db.Exec(`CREATE TABLE counter (id INTEGER PRIMARY KEY, val INTEGER)`)
	db.Exec(`INSERT INTO counter VALUES (1, 0)`)

	var wg sync.WaitGroup
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < increments; j++ {
				db.Exec(`UPDATE counter SET val = val + 1 WHERE id = 1`)
			}
		}()
	}
	wg.Wait()

	var result int
	db.QueryRow(`SELECT val FROM counter WHERE id = 1`).Scan(&result)
	return result
}

func verdict(got, want int) string {
	if got == want {
		return "✓ CORRECT"
	}
	return "✗ CORRUPTED"
}

func main() {
	fmt.Printf("=== Day 3 Probe — Why Databases Exist ===\n")
	fmt.Printf("%d goroutines × %d increments = %d expected\n\n", goroutines, increments, expected)

	file := fileCounter()
	fmt.Printf("File-based counter:  %4d / %d  %s\n", file, expected, verdict(file, expected))

	sqlite := sqliteCounter()
	fmt.Printf("SQLite counter:      %4d / %d  %s\n", sqlite, expected, verdict(sqlite, expected))

	fmt.Println("\n[?] What does SQLite do that the file doesn't?")
}
