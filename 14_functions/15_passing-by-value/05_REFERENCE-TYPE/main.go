package main

import "fmt"

func main() {
	m := make(map[string]int)
	changeMe(m)
	fmt.Println(m["Todd"]) // 44
}

func changeMe(z map[string]int) {
	z["Todd"] = 44
}

/*
Allocation with make
Back to allocation. The built-in function
serves a purpose different from new(t). I
and it returns an initialized (not zeroed)
*/
