package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for _, saluatation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(saluatation string) {
			defer wg.Done()
			fmt.Println(saluatation)
		}(saluatation)
	}
	wg.Wait()
}
