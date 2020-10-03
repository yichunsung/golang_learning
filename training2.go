package main

import "fmt"

func main() {
	// var x int

	// fmt.Println(3, "hello", false)
	// fmt.Scanln()
	// fmt.Print("請輸入一個數字：")
	// fmt.Scanln(&x)
	// fmt.Println(x)
	// 
	
	var x int
	var y int 
	fmt.Print("請輸入第一個數字：")
	fmt.Scanln(&x)
	fmt.Print("請輸入第二個數字：")
	fmt.Scanln(&y)

	var result int = x + y
	fmt.Println(result)

	var k int
	var g int

	fmt.Print("請輸入兩個數字，用空格隔開：")
	fmt.Scanln(&k, &g)

	var answer int = k + g + 3

	fmt.Println(answer)
}