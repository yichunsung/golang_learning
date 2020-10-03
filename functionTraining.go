package main

import "fmt"

func main() {
	var name string = "Tom"
	sayHelloToUser(name)
}

func sayHelloToUser(name string) {
	fmt.Println("Hello", name)
}