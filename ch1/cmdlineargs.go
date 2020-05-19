package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	echo1();
	echo2();
	echo3();
	echo4();
	echoWithIndex();
}

func echo1() {
	var s string
	// short variable declaration ":="
	sep := " "

	// for init; condition; post ...
	for i := 1; i < len(os.Args); i++ {
		if (i != 1) {
			s += sep
		}
		s += os.Args[i]
	}
	fmt.Println(s)
}

func echo2() {
	var s, sep string // both have zero value of ""

	// "_" is the blank identifier
	// range returns a pair of values: index, value
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	// lets us change how the strings are joined
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func echo4() {
	// unlike echo3, cannot choose delimiter
	fmt.Println(os.Args[1:])
}

func echoWithIndex() {
	var s, sep string
	for idx, arg := range os.Args[1:] {
		s += sep + fmt.Sprintf("[%d:%s]", idx + 1, arg)
		sep = " "
	}
	fmt.Println(s)
}
