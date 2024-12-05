package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	go generatePrimes(2, 100, ch)
	for primeNo := range ch {
		fmt.Println("Prime No :", primeNo)
	}
}

func generatePrimes(start, end int, ch chan int) {
	wg := sync.WaitGroup{}
	for no := start; no <= end; no++ {
		wg.Add(1)
		go func(no int) {
			defer wg.Done()
			if isPrime(no) {
				// time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
				ch <- no
			}
		}(no)
	}
	wg.Wait()
	close(ch)
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
