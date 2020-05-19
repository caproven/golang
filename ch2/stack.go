package main

import "fmt"

func main() {
	// test if the stack grows up or down
	i := new(int)
	fmt.Println(i)
	other()
}

func other() {
	i := new(int)
	fmt.Println(i)
}
