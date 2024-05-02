package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("Hallo ", "Print() \n")
	fmt.Println("Hallo Println()")
	var s = "Printf()"
	fmt.Printf("Hallo %s \n", s)

	var nr = 42
	var name = "Obiwan Kenobi"
	fmt.Printf("%03d: Hallo, %s\n", nr, name)
	s = fmt.Sprintf("%03d: Hallo, %s\n", nr, name) //
	fmt.Print(s)

	file, _ := os.Create("datei.txt")
	fmt.Fprintf(file, "%03d: Hallo, %s\n", nr, name)
	file.Close()
}

// Hallo Print()
//Hallo Println()
// Hallo Printf()
// 042: Hallo, Obiwan Kenobi
