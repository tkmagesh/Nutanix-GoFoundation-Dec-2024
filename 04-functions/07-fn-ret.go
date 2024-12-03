package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for range 20 {
		time.Sleep(500 * time.Millisecond)
		fn := getFn()
		fn()
	}
}

func getFn() func() {
	if randNo := rand.Intn(2000); randNo%2 == 0 {
		return f1
	}
	return f2
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
