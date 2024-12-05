package main

import (
	"fmt"
	"sync"
)

// Custom type encapsulating the data and concurrent safe logic for manipulating the data
type SafeCounter struct {
	sync.Mutex //composition
	count      int
}

func (sf *SafeCounter) Add(delta int) {
	sf.Lock()
	{
		sf.count++
	}
	sf.Unlock()
}

var sf SafeCounter

func main() {
	wg := &sync.WaitGroup{}
	for range 200 {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("count :", sf.count)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	sf.Add(1)
}
