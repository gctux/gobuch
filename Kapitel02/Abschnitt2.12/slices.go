package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%p - len: %d cap: %d %#v\n", s, len(s), cap(s), s)
}

func main() {
	var a []int
	printSlice(a)
	a = append(a, 1)
	printSlice(a)
	a = append(a, 2)
	printSlice(a)
	a = append(a, 3)
	printSlice(a)
}

// Output: 1000
// 0xc0000120d8 - len: 1 cap: 1 []int{1}
// 0xc000012100 - len: 2 cap: 2 []int{1, 2}
// 0xc000022160 - len: 3 cap: 4 []int{1, 2, 3}
