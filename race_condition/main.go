package main

import (
	"fmt"
	"sync"
	"time"
)

// *sync.Mutex
// type counter struct {
// 	count int
// 	mu    *sync.Mutex
// }

type counter struct {
	count int
	mu    *sync.RWMutex
}

func (c *counter) inc() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// *sync.Mutex
// func (c *counter) value() int {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	return c.count
// }

func (c *counter) value() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}

func main() {
	// #1 WARNING: DATA RACE
	// counter := 0
	// for i := 0; i < 1000; i++ {
	// 	go func() {
	// 		counter++
	// 	}()
	// }
	// time.Sleep(time.Second)

	// fmt.Println(counter)

	// #2 Safe code
	c := counter{
		mu: new(sync.RWMutex),
	}
	for i := 0; i < 1000; i++ {
		go func() {
			c.inc()
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(c.value())
}
