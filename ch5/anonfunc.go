package main

import "fmt"

func performOperation(x, y int, operation func(a, b int) int) int {
	return operation(x, y)
}

func main() {
	result := performOperation(6, 8, func(a, b int) int {
		return a * b
	})
	fmt.Println(result) // "48"
}
