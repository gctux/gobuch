package main

import "fmt"

func main() {
	fmt.Print("Hallo ", "Print()\n")
	fmt.Println("Hello Println()")
	var s = "Printf()"
	fmt.Printf("Hallo, %s\n", s)
}

//Output
//Hallo Print()
//Hello Println()
//Hallo, Printf()
