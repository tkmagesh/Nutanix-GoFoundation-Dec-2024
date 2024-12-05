package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1) // increment the counter by 1
	go f1(wg) //sheduling the execution of f1() through the scheduler
	f2()
	// blocking the execution of main() so that the scheduler can look for other goroutines that are scheduled and execute them

	wg.Wait() //block the execution of this function (main()) until the counter becomes 0 (default = 0)
}

func f1(wg *sync.WaitGroup) {
	defer wg.Done() // decrementing the counter by 1
	fmt.Println("f1 started")
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
