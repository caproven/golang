package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	for i, r := range s {
		fmt.Printf("Idx: %d\tRune: %q\n", i, r)
	}

	fmt.Println(string(1234567)) // prints the Unicode replacemnet character
}
