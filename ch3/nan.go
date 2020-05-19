package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.NaN())
	fmt.Println(math.Sqrt(-1))
	fmt.Printf("%T\n", math.NaN())
	fmt.Printf("%t\n", 5 != math.NaN())
}
