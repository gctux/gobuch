package main

import "fmt"

func main() {
	type äpfel int
	type birnen int
	a := äpfel(10)
	b := birnen(10)
	anzahl := int(a) + int(b)
	fmt.Printf("Anzahl Früchte: %d", anzahl)
}

// Output:
// Anzahl Früchte: 20
