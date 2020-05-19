package main

import "fmt"

func main() {
	list := []int{1,2,3,4,5}
	fmt.Println(sumInts(list))
}

func sumInts(nums []int) (sum int) {
	for _, v := range nums {
		sum += v
	}
	return sum
}
