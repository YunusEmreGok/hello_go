package main

import "fmt"

var first float64
var second float64

func main() {
	fmt.Print("Enter First Number.")
	fmt.Scan(&first)
	fmt.Print("Enter Second Number. ")
	fmt.Scan(&second)

	fmt.Print("Conclusion : ", second+first, "\n")

	fmt.Print("Multiply  : ", second*first, "\n")

	fmt.Print("Divide  : ", second/first, "\n")
}
