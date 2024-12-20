package main

import (
	"fmt"
	"sync"
)

// communicate by sharing memory (Go does not advise this)
var result int

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(100, 200, wg)
	wg.Wait()
	fmt.Println(result)
}

func add(x, y int, wg *sync.WaitGroup) {
	defer wg.Done()
	result = x + y
}
