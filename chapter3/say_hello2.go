package main

import (
	"fmt"
	"sync"
)

func sayHello() {
	defer wg.Done()
	fmt.Println("hello")
}

var wg sync.WaitGroup = sync.WaitGroup(0, 0)

func main() {
	wg.Add(1)
	go sayHello()
	wg.Wait()
}
