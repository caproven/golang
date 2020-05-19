package main

import "fmt"

func main() {
	x := 5
	myPointer := &x
	fmt.Printf("Type: %T\tAddr: %v\tValue: %d\n", myPointer, myPointer, *myPointer)

	intPtr1 := new(int)
	var myInt int
	intPtr2 := &myInt
	fmt.Printf("%T %d %T %d\n", intPtr1, *intPtr1, intPtr2, *intPtr2)
}
