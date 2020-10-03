package main
import "fmt"


func main() {
	var msg string = "----> Script <----"
	var x int

	for x = 0; x < 10; x+=2 {
		fmt.Println(x)
	}
	fmt.Println(msg);

	var result int = 0
	var i int

	for i = 0; i <= 50; i++ {
		result = result + i
	}
	fmt.Println(result)
	// break line
	fmt.Println(msg)

	for i = 0; i < 5; i++ {
		if i == 2 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println(msg)

	for i = 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println(msg)
}

