package main

import "fmt"

func main() {
	var a int
	var b *int // Typ Pointer auf int
	a = 123
	b = &a // Speicheradresse mit &
	fmt.Println(b, *b)
	*b = 100 //Dereferenzierung
	fmt.Println(a)
}

// Output:
// 0xc0000120d0 123
// 100
