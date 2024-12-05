package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() //sheduling the execution of f1() through the scheduler
	f2()
	// blocking the execution of main() so that the scheduler can look for other goroutines that are scheduled and execute them

	// DO NOT sync goroutines using the following approaches (these are poor man' sync techniques)
	// time.Sleep(time.Millisecond)
	// fmt.Scanln()

	time.Sleep(2 * time.Second)
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(4 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
