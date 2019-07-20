package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i, saluatation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			fmt.Println(i)
			defer wg.Done()
			fmt.Println(saluatation)
		}()
	}
	wg.Wait()
}
