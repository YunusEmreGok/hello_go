package main

import "fmt"

var first int
var second int

func main() {
	fmt.Print("Enter First Number.")
	fmt.Scan(&first)
	fmt.Print("Enter Second Number. ")
	fmt.Scan(&second)

	fmt.Print("Conclusion : ", second+first)
}
