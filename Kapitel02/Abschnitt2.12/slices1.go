package main

import "fmt"

func printSlice(name string, slice ...[]int) {
	fmt.Println(name, ":")
	for _, s := range slice {
		fmt.Printf("%p - len: %d cap: %d - %#v\n", s, len(s), cap(s), s)
	}
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := a[1:5]
	printSlice("b := a[1:5]", a, b)
	b = append(b, 11)
	printSlice("nach append", a, b)
	c := a[1:5:5]
	printSlice("c := a[1:5:5]", c)
	c = append(c, 14)
	printSlice("nach c append", a, c)
}

// Output:
// b := a[1:5] :
// 0xc0000240f0 - len: 10 cap: 10 - []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
// 0xc0000240f8 - len: 4 cap: 9 - []int{1, 2, 3, 4}
// nach append :
// 0xc0000240f0 - len: 10 cap: 10 - []int{0, 1, 2, 3, 4, 11, 6, 7, 8, 9}
// 0xc0000240f8 - len: 5 cap: 9 - []int{1, 2, 3, 4, 11}
// c := a[1:5:5] :
// 0xc0000240f8 - len: 4 cap: 4 - []int{1, 2, 3, 4}
// nach c append :
// 0xc0000240f0 - len: 10 cap: 10 - []int{0, 1, 2, 3, 4, 11, 6, 7, 8, 9}
// 0xc0000201c0 - len: 5 cap: 8 - []int{1, 2, 3, 4, 14}
