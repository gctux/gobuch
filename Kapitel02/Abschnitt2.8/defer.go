package main

import "fmt"

func meineFunktion() {
	defer func() {
		fmt.Println(1)
		fmt.Println(2)
	}()
	fmt.Println(3)
}

func main() {
	meineFunktion()
}

// Output:
// 3
// 1
// 2
