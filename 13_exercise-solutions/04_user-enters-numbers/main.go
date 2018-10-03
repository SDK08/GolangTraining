package main

import "fmt"

func main() {
	var numOne int
	var numTwo int
	fmt.Print("Please enter your first number: ")
	fmt.Scan(&numOne)
	fmt.Print("Please enter your second number: ")
	fmt.Scan(&numTwo)
	fmt.Println(numOne, " / ", numTwo, " = ", numOne/numTwo)
}
