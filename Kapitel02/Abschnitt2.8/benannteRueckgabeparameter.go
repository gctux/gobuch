package main

import "fmt"

func foo(s string) (a string, b string) {
	b = s
	return
}

func main() {
	a, b := foo("Hallo Lutz!")
	fmt.Println(a)
	fmt.Println(b)
}

// Output:
//
// Hallo Lutz!
