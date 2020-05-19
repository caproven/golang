package main

import "fmt"

func main() {
	// map literal
	ages := map[string]int{
		"jeremy": 23,
		"charles": 19,
	}

	fmt.Println(ages["charles"]) // 19
	ages["bob"]++
	fmt.Println(ages["bob"])
}
