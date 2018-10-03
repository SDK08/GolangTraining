package main

import "fmt"

func main() {

	age := 44
	fmt.Println(&age) // 0xc000098000

	changeMe(&age)

	fmt.Println(&age) // 0xc000098000
	fmt.Println(age)  //24
}

func changeMe(z *int) {
	fmt.Println(z)  // 0xc000098000
	fmt.Println(*z) // 44
	*z = 24
	fmt.Println(z)  // 0xc000098000
	fmt.Println(*z) // 24
}
