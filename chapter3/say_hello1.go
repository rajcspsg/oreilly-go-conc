package main

import "fmt"

func sayHello() {
	fmt.Println("hello")
}

func main() {
	go sayHello()
}
