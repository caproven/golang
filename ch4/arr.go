package main

import "fmt"

func main() {
	arr2:= [3]int{1,2}
	fmt.Println(arr2)

	arr3 := [...]int{5: 1}
	fmt.Println(arr3)

	var arr4 [1]int
	fmt.Println("arr4", arr4)
	incIdxZero(&arr4)
	fmt.Println("arr4", arr4)
}

func incIdxZero(ptr *[1]int) {
	fmt.Printf("is type: %T\n", ptr)
	ptr[0]++
}
