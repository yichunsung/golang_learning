package main // 可執行程式必須使用main 封包
import "fmt" // 載入內建的 fmt 封包 用來做基本輸出輸入

func main() {
	// run the script
	// export to termainal
	/*
	fmt.Println(3)
	fmt.Println(3.123)
	fmt.Println("測試中文")
	fmt.Println(false)
	fmt.Println('a') // rune
	*/

	var dataX int //宣導變數
	dataX = 4 
	fmt.Println(dataX)

	var dataNumber float64
	dataNumber = 3.2123
	fmt.Println(dataNumber)
	dataNumber = 123

	var stringData string = "Hello World Golang!!!!"
	fmt.Println(stringData)
}