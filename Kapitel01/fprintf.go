package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Create("datei.txt")
	var nr = 2
	var name = "Lutz"
	fmt.Fprintf(file, "%03d: Hallo, %s\n", nr, name)
	file.Close()
}

// Output:
// 002: Hallo, Lutz
