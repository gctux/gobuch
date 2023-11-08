package main

import "fmt"

func main() {
	type äpfel int
	type birnen int
	a := äpfel(10)
	b := birnen(10)
	fmt.Println(a + b)
}

// Output:
// invalid operation: a + b (mismatched types äpfel and birnen)
