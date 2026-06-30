package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	writes        = 10
	readsPerWrite = 100
	totalReads    = writes * readsPerWrite

	sourceReadLatency  = 500 * time.Microsecond
	sourceWriteLatency = 800 * time.Microsecond
)

type Product struct {
	Name       string
	PriceCents int
	Version    int
}

type Metrics struct {
	SourceReads   int
	SourceWrites  int
	CacheHits     int
	CacheMisses   int
	Invalidations int
	Elapsed       time.Duration
}

type SourceOfTruth struct {
	mu     sync.Mutex
	data   Product
	reads  int
	writes int
}

func NewSourceOfTruth() *SourceOfTruth {
	return &SourceOfTruth{
		data: Product{
			Name:       "system-design-notebook",
			PriceCents: 2500,
			Version:    0,
		},
	}
}

func (s *SourceOfTruth) Read() Product {
	time.Sleep(sourceReadLatency)

	s.mu.Lock()
	defer s.mu.Unlock()

	s.reads++
	return s.data
}

func (s *SourceOfTruth) Write(priceCents int) {
	time.Sleep(sourceWriteLatency)

	s.mu.Lock()
	defer s.mu.Unlock()

	s.data.PriceCents = priceCents
	s.data.Version++
	s.writes++
}

func (s *SourceOfTruth) Metrics() Metrics {
	s.mu.Lock()
	defer s.mu.Unlock()

	return Metrics{
		SourceReads:  s.reads,
		SourceWrites: s.writes,
	}
}

type ReadThroughCache struct {
	source *SourceOfTruth

	mu            sync.Mutex
	cached        Product
	hasCached     bool
	hits          int
	misses        int
	invalidations int
}

func NewReadThroughCache(source *SourceOfTruth) *ReadThroughCache {
	return &ReadThroughCache{source: source}
}

func (c *ReadThroughCache) Read() Product {
	c.mu.Lock()
	if c.hasCached {
		c.hits++
		product := c.cached
		c.mu.Unlock()
		return product
	}
	c.misses++
	c.mu.Unlock()

	product := c.source.Read()

	c.mu.Lock()
	c.cached = product
	c.hasCached = true
	c.mu.Unlock()

	return product
}

func (c *ReadThroughCache) Write(priceCents int) {
	c.source.Write(priceCents)

	c.mu.Lock()
	defer c.mu.Unlock()

	c.hasCached = false
	c.invalidations++
}

func (c *ReadThroughCache) Metrics() Metrics {
	c.mu.Lock()
	defer c.mu.Unlock()

	return Metrics{
		CacheHits:     c.hits,
		CacheMisses:   c.misses,
		Invalidations: c.invalidations,
	}
}

func runDirectReads() Metrics {
	source := NewSourceOfTruth()
	start := time.Now()

	for write := 1; write <= writes; write++ {
		source.Write(2500 + write)
		for read := 0; read < readsPerWrite; read++ {
			_ = source.Read()
		}
	}

	metrics := source.Metrics()
	metrics.Elapsed = time.Since(start)
	return metrics
}

func runReadOptimized() Metrics {
	source := NewSourceOfTruth()
	cache := NewReadThroughCache(source)
	start := time.Now()

	for write := 1; write <= writes; write++ {
		cache.Write(2500 + write)
		for read := 0; read < readsPerWrite; read++ {
			_ = cache.Read()
		}
	}

	sourceMetrics := source.Metrics()
	cacheMetrics := cache.Metrics()

	cacheMetrics.SourceReads = sourceMetrics.SourceReads
	cacheMetrics.SourceWrites = sourceMetrics.SourceWrites
	cacheMetrics.Elapsed = time.Since(start)
	return cacheMetrics
}

func printMetrics(label string, metrics Metrics) {
	fmt.Printf("%s:\n", label)
	fmt.Printf("  source reads:   %4d\n", metrics.SourceReads)
	fmt.Printf("  source writes:  %4d\n", metrics.SourceWrites)
	if metrics.CacheHits+metrics.CacheMisses > 0 {
		fmt.Printf("  cache hits:     %4d\n", metrics.CacheHits)
		fmt.Printf("  cache misses:   %4d\n", metrics.CacheMisses)
		fmt.Printf("  invalidations:  %4d\n", metrics.Invalidations)
	}
	fmt.Printf("  elapsed:        %s\n\n", metrics.Elapsed.Round(time.Millisecond))
}

func main() {
	fmt.Println("=== Day 4 Probe - Read vs. Write Asymmetry ===")
	fmt.Printf("Workload: %d reads, %d writes = %d:1 read/write ratio\n\n", totalReads, writes, readsPerWrite)

	direct := runDirectReads()
	printMetrics("Direct source-of-truth reads", direct)

	optimized := runReadOptimized()
	printMetrics("Read-optimized path", optimized)

	fmt.Println("[?] If reads dominate, which path deserves the most design attention?")
	fmt.Println("[!] But every write now has an extra job: keep derived read state correct.")
}
