package main

import "fmt"

func main() {
	numbers := []int{0,1,2,3,4,5,6,7,8,9}
	rev(numbers)
	fmt.Println(numbers)
}

// reverse a slice of ints in place
func rev(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
