package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dgraph-io/ristretto"
)

func main() {
	cache, err := ristretto.NewCache(&ristretto.Config{

		// 10x the number of items you expect to keep in the cache when full
		NumCounters: 1e7, // number of keys to track frequency of (10M).

		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
		Metrics:     true,
	})
	if err != nil {
		log.Fatal(err)
	}

	// set a value with a cost of 1
	// the cost is a value that represent the "weight" of the cache
	// a heavy cache should have a bigger value
	// when the cache is full the bigger value will be dropped
	cache.Set("key", "value", 1)
	fmt.Println("Hits", cache.Metrics.Hits())

	// wait for value to pass through buffers
	time.Sleep(10 * time.Millisecond)

	fmt.Println("CostAdded", cache.Metrics.CostAdded())

	value, found := cache.Get("key")
	if !found {
		log.Fatal("missing value")
	}
	fmt.Println(value)
	fmt.Println("CostAdded", cache.Metrics.CostAdded())

	fmt.Println("Hits", cache.Metrics.Hits())

	fmt.Println("Misses", cache.Metrics.Misses())
	cache.Get("key1")
	fmt.Println("Misses", cache.Metrics.Misses())

	fmt.Println("CostAdded", cache.Metrics.CostAdded())
	cache.Del("key")

	fmt.Println(cache.Metrics.String())
}
