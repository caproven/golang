// since embedded fields let us shortcut to underlying fields,
// what happens if the underlying fields overlap? Or if there
// are duplicates between embedded fields and explicit fields?

// FINDINGS
// if underlying fields and explicit fields share the same name,
// shorcutting does not work and priority is given to the explicit field.



package main

import "fmt"

type Point struct {
	X, Y int
}

type Point2 struct {
	X, Y int
}

type Circle struct {
	Point
	X, Y int // maybe won't let compile?
}

type Square struct {
	Point
	Point2
}

func main() {
	c := Circle{
		Point: Point{1,2},
		X: 3,
		Y: 4,
	}
	fmt.Printf("%#v\n", c)
	fmt.Println(c.X) // "3" -> explicit field used instead of shortcutting to c.Point.x

	s := Square{
		Point: Point{1,2},
		Point2: Point2{3,4},
	}
	fmt.Printf("%#v\n", s)
	fmt.Println(s.Y) // -> gives "ambiguous selector" compiler error
}
