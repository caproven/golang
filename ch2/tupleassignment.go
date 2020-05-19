package main

import "fmt"

func main() {
	x, y := 1, 2
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)
}
