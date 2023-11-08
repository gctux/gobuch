package main

import "fmt"

func foo() *int {
	bar := 123
	return &bar
}

func main() {
	a := foo()
	fmt.Println(a, *a)
}

// Output:
// 0xc0000120d0 123
