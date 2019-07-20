package main

import (
	"fmt"
	"sync"
)

func main() {
	var memAccess sync.Mutex
	var value int

	go func() {
		memAccess.Lock()
		value++
		memAccess.Unlock()
	}()

	memAccess.Lock()
	if value == 0 {
		fmt.Println(value)
	} else {
		fmt.Printf("the value is %v.\n", value)
	}
	memAccess.Unlock()
}
