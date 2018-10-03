package main

import "fmt"

func main() {

	if false {
		fmt.Println("first print statement")
	} else if false {
		fmt.Println("second print statement")
	} else if true {
		fmt.Println("third times a charm")
	} else {
		fmt.Println("Not getting to this print statement")
	}
}
