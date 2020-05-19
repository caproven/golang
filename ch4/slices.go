package main

import "fmt"

func main() {
	days := []string{"m", "tu", "w", "th", "f", "sa", "su"}
	weekDays := days[:5]
	fmt.Println("weekdays:", weekDays)
	fmt.Println("weekend:", days[5:])
	fmt.Println("all days:", weekDays[:7])

	fmt.Printf("type of weekdays: %T\n", weekDays)
	testNil()
	testAppend()
}

// assert that len(s) & cap(s) return 0 if s is nil
func testNil() {
	var nilSlice []int
	fmt.Println("nilSlice:", nilSlice)
	fmt.Println("len:", len(nilSlice))
	fmt.Println("cap:", cap(nilSlice))
	// can you append to a nil slice?
	nilSlice = append(nilSlice, 2)
	fmt.Println("nilSlice:", nilSlice)
}

func testAppend() {
	printSlice := func (s []rune) {
		fmt.Printf("%q\t%d\t%d\n", string(s), len(s), cap(s))
	}

	var runes []rune
	newRunes := append(runes, 'H')
	fmt.Println("slice\tlen\tcap")
	printSlice(runes)
	printSlice(newRunes)

	fmt.Println("--")

	newRunes2 := append(newRunes, 'I')
	printSlice(newRunes)
	printSlice(newRunes2)
}
