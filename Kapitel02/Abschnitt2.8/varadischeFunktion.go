package main

import "fmt"

func add(a ...int) int {
	summe := 0
	for _, v := range a {
		summe = summe + v
	}
	return summe
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add())
	fmt.Println(add(1, 2, 4, 5))
}

// Output:
// 3
// 0
// 12
