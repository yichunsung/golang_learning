package main
import "fmt"

func main() {
	
	for true {
		var input int
		fmt.Scanln(&input)
		if input == 0 {
			break
		}
		fmt.Println(input)
		fmt.Println(input + 2)
	}
}