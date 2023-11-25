package main

import "fmt"

func main() {

	a := 10
	switch {
	case a > 3:
		fmt.Println("a > 3")
		fallthrough
	case a > 5:
		fmt.Println("a > 5")
	case a > 8:
		fmt.Println("a > 8")
	case a > 15:
		fmt.Println("a > 15")
	}
}
