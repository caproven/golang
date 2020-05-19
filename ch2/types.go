package main

import "fmt"

type Feet float64
type Meters float64

const (
	HeightF Feet = 5.6
	Wingspan Feet = 5.1
)

func main() {
	var testMeters Meters = 4.4
	fmt.Println(HeightF + testMeters)
}
