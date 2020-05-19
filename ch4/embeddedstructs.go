package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point // embedded
	Radius int
}

func main() {
	c := Circle{
		Point: Point{5,7},
		Radius: 3,
	}

	fmt.Println(c.X) // "5"
	fmt.Println(c.Point.Y) // "7"

	fmt.Printf("%#v\n", c)
}
