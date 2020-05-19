package main

import (
	"fmt"
	"strings"
)

const substring = "$foo"

func main() {
	input := "$foosome $foo$foo$foostring with $foo"
	output := expand(input, func(s string) string {
		return s + "BAR"
	})
	fmt.Println(output)
}

// TODO - redo using byte.Buffer or something
func expand(s string, f func(string) string) string {
	fmt.Println("->", s)
	i := strings.Index(s, substring)
	if (i == -1) {
		return s
	}
	return s[:i] + f(substring) + expand(s[i+len(substring):], f)
}
