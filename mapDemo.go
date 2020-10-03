package main

import "fmt"

func main() {
	// var m map [int] int

	// m = make(map[int] int)
	// fmt.Print(m)

	// s := map [string] string { "abc": "Hellp", "gogog": "world" }

	// fmt.Println(s)

	type student struct {
		name string
		age int
	}

	fmt.Println(student{ "tom", 12 })

	var sdd = student { "hello", 33 }

	fmt.Print(sdd.name)
}