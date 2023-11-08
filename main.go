package main

import "fmt"

func main() {
	var nr = 2
	var name = "Lutz"
	s := fmt.Sprintf("%03d: Hallo, %s\n", nr, name)
	fmt.Print(s)
}

// Output:
// 002: Hallo, Lutz
