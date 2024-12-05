package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1) // increment the counter by 1
	go f1()   //sheduling the execution of f1() through the scheduler
	f2()
	// blocking the execution of main() so that the scheduler can look for other goroutines that are scheduled and execute them

	wg.Wait() //block the execution of this function (main()) until the counter becomes 0 (default = 0)
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(4 * time.Second)
	fmt.Println("f1 completed")
	wg.Done() // decrementing the counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
