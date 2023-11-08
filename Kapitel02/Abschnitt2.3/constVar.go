package main

import "fmt"

func add2(a int) int {
	return 2 + a
}

func main() {
	const c = 2.0
	fmt.Println((add2(c)))
	//4
	var v = 2.0
	fmt.Println((add2(v)))
	//cannot use v (variable of type float64) as int value in argument to add2
}
