package main

import (
	"flag"
	"fmt"
)

func main() {
	var country = flag.String("country", "United States", "user's country of origin")
	flag.Parse()

	fmt.Println("Got the country flag:", *country)
}
