package main

import "fmt"

func main() {

	a := 0

	// Klassisches switch
	switch a {
	case 0:
		fmt.Println("a = 0")
	case 2, 3, 4:
		fmt.Println("2, 3 oder 4")
	default:
		fmt.Println("alles andere")
	}

	// Switch mit Deklaration
	// b wird per Kurzdeklaration für den Block definiert
	switch b := a - 1; b {
	case 0:
		fmt.Println("b = 0")
	default:
		fmt.Println("alles andere")
	}
	// b ist nicht mehr gültig

	// kein Ausdruck nach switch
	switch {
	case a == 0:
		fmt.Println("a = 0")
	default:
		fmt.Println("alles andere")
	}

}
