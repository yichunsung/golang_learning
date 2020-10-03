package main

import "fmt"

func main() {
	var money int

	fmt.Println("請問想領多少錢？")
	fmt.Scanln(&money)

	if money < 100 {
		fmt.Println("too few")
	} else if money < 1000 {
		fmt.Println("ok")
	} else {
		fmt.Println("too much")
	}
}