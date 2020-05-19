package main

import (
	"bufio"
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

func main() {
	dup3()
}

func dup1() {
	counts := make(map[string]int)
	// input is type bufio.Scanner
	input := bufio.NewScanner(os.Stdin)
	// while there is more stuff to scan..
	for input.Scan() {
		// if key isn't found in the map, it returns
		// the zero value for the value (in this case, int->0)
		//
		// we know there is more stuff to scan,
		// so go read it as text
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// reads from stdin or a list of files in Args
func dup2() {
	counts := make(map[string]int)
	fileNames := os.Args[1:]
	if len(fileNames) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, fileName := range fileNames {
			file, err := os.Open(fileName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(file, counts)
			file.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

func dup3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		// data is a byte slice and must later be converted to a string
		data, err := ioutil.ReadFile(filename)
		if err != nil { // error occurred
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

