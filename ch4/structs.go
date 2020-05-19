package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main() {
	type MyStruct struct {
		Field1 int
		Field2 int
		Field3 rune
		Field4 string
	}

	var s1 MyStruct
	fmt.Println(s1) // {0 0 0  }
	fmt.Printf("%d\n", s1.Field3)
	fmt.Printf("%t\n", s1.Field3 == 0)

	testStructPtr()
}

func testStructPtr() {
	me := Person{"my name", 21}
	mePtr := &me
	fmt.Println(me.Age) // "21"
	fmt.Println(mePtr.Age) // "21"
	fmt.Println((*mePtr).Age) // "21"

	fmt.Printf("me is %T\n", me)
	fmt.Printf("mePtr is %T\n", mePtr)
}
