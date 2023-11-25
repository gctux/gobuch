package main

import "fmt"

func getNumber() int {
	return 11
}

func main() {
	x := 4
	if x > 5 {
		fmt.Printf("%d ist größer 5\n", x)
	} else if x == 5 {
		fmt.Printf("%d ist gleich 5\n", x)
	} else {
		fmt.Printf("%d ist kleiner 5\n", x)
	}

	if a := getNumber(); a > 10 {
		fmt.Println("a > 10")
	} else {
		fmt.Println(a)
	}

	// ab hier existiert a nicht mehr

}
